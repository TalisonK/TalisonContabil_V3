package model

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFullIncomes() ([]domain.IncomeDTO, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if !statusDBLocal {
		incomes := []domain.Income{}
		err := database.DBlocal.Find(&incomes).Error

		if err != nil {
			e := logging.FailedToFindOnDB("All Incomes", "Incomes", nil)
			return nil, util.GetTagError(http.StatusInternalServerError, fmt.Errorf(e))
		}

		incomeDto := []domain.IncomeDTO{}

		for i := range incomes {
			incomeDto = append(incomeDto, incomes[i].ToDTO())
		}

		logging.FoundOnDB("All Incomes", constants.LOCAL)
		return incomeDto, nil
	}

	if statusDBCloud {

		filter := bson.M{}

		cursor, err := database.DBCloud.Income.Find(context.Background(), filter)

		if err != nil {
			e := logging.FailedToFindOnDB("All Incomes", "Incomes", err)
			return nil, util.GetTagError(http.StatusInternalServerError, fmt.Errorf(e))
		}
		incomes := []domain.IncomeDTO{}

		for cursor.Next(context.Background()) {
			var income primitive.M
			cursor.Decode(&income)

			incomes = append(incomes, domain.PrimToIncome(income).ToDTO())
		}

		logging.FoundOnDB("All Incomes", constants.CLOUD)

		return incomes, nil

	}
	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func GetUserIncomes(id string) ([]domain.IncomeDTO, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDBLocal {
		incomes := []domain.Income{}

		result := database.DBlocal.Where("user_id = ?", id).Find(&incomes)

		if result.Error != nil {
			logging.FailedToFindOnDB(id, constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		}

		incomesDto := []domain.IncomeDTO{}

		for _, income := range incomes {
			incomesDto = append(incomesDto, income.ToDTO())
		}

		logging.FoundOnDB(id, constants.LOCAL)
		return incomesDto, nil
	}

	if statusDBCloud {
		incomes := []domain.IncomeDTO{}

		user := bson.M{"_id": id}
		filter := bson.M{"user": user}

		cursor, err := database.DBCloud.Income.Find(context.Background(), filter)

		if err != nil {
			logging.FailedToFindOnDB(id, constants.CLOUD, err)
			return nil, util.GetTagError(http.StatusBadRequest, err)
		}

		for cursor.Next(context.Background()) {
			var aux bson.M
			cursor.Decode(aux)
			incomes = append(incomes, domain.PrimToIncome(aux).ToDTO())
		}

		logging.FoundOnDB(id, constants.CLOUD)
		return incomes, nil
	}
	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func GetIncomesByDate(userId string, startingDate string, endingDate string) ([]domain.Income, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	incomes := []domain.Income{}

	if statusDBLocal {

		result := database.DBlocal.Where("User_id = ? AND received_at between ? AND ?", userId, startingDate, endingDate).Find(&incomes)

		if result.Error != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Incomes from user %s", userId), constants.LOCAL, result.Error)
		}

		logging.FoundOnDB(fmt.Sprintf("Incomes from user %s", userId), constants.CLOUD)
		return incomes, nil
	}

	if statusDBCloud {

		auxSD, _ := time.Parse(time.RFC3339, startingDate)
		sd := primitive.NewDateTimeFromTime(auxSD)

		auxED, _ := time.Parse(time.RFC3339, endingDate)
		ed := primitive.NewDateTimeFromTime(auxED)

		sdBson := bson.M{"$gt": sd, "$lt": ed}
		filter := bson.M{"userID": userId, "receivedAt": sdBson}

		cursor, err := database.DBCloud.Income.Find(context.Background(), filter)

		if err != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Incomes from user %s", userId), constants.CLOUD, err)
		}

		for cursor.Next(context.Background()) {
			var aux bson.M

			cursor.Decode(aux)

			incomes = append(incomes, *domain.PrimToIncome(aux))
		}

		logging.FoundOnDB(fmt.Sprintf("Incomes from user %s", userId), constants.CLOUD)
		return incomes, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func CreateIncome(income domain.IncomeDTO) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if income.Value == 0 || income.Description == "" || income.ReceivedAt == "" || income.UserID == "" {
		return util.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.InvalidFields()))
	}

	income.CreatedAt = util.GetTimeNow()
	income.UpdatedAt = util.GetTimeNow()

	if statusDBCloud {

		entity := income.ToEntity()

		raw := entity.ToPrim()

		result, err := database.DBCloud.Income.InsertOne(context.Background(), raw)

		if err != nil {
			logging.FailedToCreateOnDB(income.Description, constants.CLOUD, err)
			return util.GetTagError(http.StatusBadRequest, err)
		}

		income.ID = result.InsertedID.(primitive.ObjectID).Hex()

		logging.CreatedOnDB(income.ID, constants.CLOUD)
	}

	if statusDBLocal && income.ID != "" {

		entity := income.ToEntity()

		result := database.DBlocal.Create(&entity)

		if result.Error != nil {
			logging.FailedToCreateOnDB(income.ID, constants.LOCAL, result.Error)
			return util.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.CreatedOnDB(income.ID, constants.LOCAL)
		return nil
	}

	return util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func UpdateIncome(income domain.IncomeDTO) (*domain.IncomeDTO, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	baseIncome, err := findIncomeById(income.ID)

	if err != nil {
		logging.FailedToFindOnDB(income.ID, constants.LOCAL, err.Inner)
		return nil, err
	}

	baseIncome.UpdatedAt = util.GetTimeNow()

	if income.Description != "" {
		baseIncome.Description = income.Description
	}

	if income.ReceivedAt != "" {
		baseIncome.ReceivedAt = income.ReceivedAt
	}

	if income.Value != 0 {
		baseIncome.Value = income.Value
	}

	if statusDBLocal {

		result := database.DBlocal.Save(baseIncome)

		if result.Error != nil {
			logging.FailedToUpdateOnDB(baseIncome.ID, constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		}
		logging.UpdatedOnDB(baseIncome.ID, constants.LOCAL)
	}

	if statusDBCloud {

		incomeParse := baseIncome.ToPrim()

		id, _ := primitive.ObjectIDFromHex(income.ID)

		filter := bson.M{"_id": id}

		_, err := database.DBCloud.Income.ReplaceOne(context.Background(), filter, incomeParse)

		if err != nil {
			logging.FailedToUpdateOnDB(baseIncome.ID, constants.CLOUD, err)
			return nil, util.GetTagError(http.StatusBadRequest, err)
		}

		logging.UpdatedOnDB(baseIncome.ID, constants.CLOUD)
		dto := baseIncome.ToDTO()
		return &dto, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func DeleteIncome(id string) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDBLocal {

		result := database.DBlocal.Delete(&domain.Income{}, "id = ?", id)

		if result.Error != nil {
			logging.FailedToDeleteOnDB(id, constants.LOCAL, result.Error)
			return util.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.DeletedOnDB(id, constants.LOCAL)
	}

	if statusDBCloud {

		idPrim, _ := primitive.ObjectIDFromHex(id)

		filter := bson.M{"_id": idPrim}

		_, err := database.DBCloud.Income.DeleteOne(context.Background(), filter)

		if err != nil {
			logging.FailedToDeleteOnDB(id, constants.CLOUD, err)
			return util.GetTagError(http.StatusBadRequest, err)
		}

		logging.DeletedOnDB(id, constants.CLOUD)
	}

	return nil

}

func findIncomeById(id string) (*domain.Income, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	var income domain.Income

	if statusDBLocal {

		result := database.DBlocal.Find(&income, "id = ?", id)

		if result.Error != nil {
			logging.FailedToFindOnDB(id, constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.FoundOnDB(id, constants.LOCAL)
		return &income, nil
	}

	if statusDBCloud {

		raw := bson.M{}

		filter := bson.M{"_id": id}

		cursor, err := database.DBCloud.Income.Find(context.Background(), filter)

		if err != nil {
			logging.FailedToFindOnDB(id, constants.LOCAL, err)
			return nil, util.GetTagError(http.StatusBadRequest, err)
		}

		cursor.Decode(raw)

		dto := domain.PrimToIncome(raw)

		return dto, nil

	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())

}

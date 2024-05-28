package model

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFullIncomes() ([]domain.IncomeDTO, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.GetIncomes"))
	}

	if !statusDBLocal {
		incomes := []domain.IncomeDTO{}
		err := database.DBlocal.Find(&incomes).Error

		if err != nil {
			e := logging.FailedToFindOnDB("All Incomes", "Incomes", nil, "model.GetIncomes")
			return nil, util.GetTagError(http.StatusInternalServerError, fmt.Errorf(e))
		}

		logging.FoundOnDB("All Incomes", "Local", "model.GetIncomes")
		return incomes, nil
	}

	if statusDBCloud {

		filter := bson.M{}

		cursor, err := database.DBCloud.Income.Find(context.Background(), filter)

		if err != nil {
			e := logging.FailedToFindOnDB("All Incomes", "Incomes", err, "model.GetIncomes")
			return nil, util.GetTagError(http.StatusInternalServerError, fmt.Errorf(e))
		}
		incomes := []domain.IncomeDTO{}

		for cursor.Next(context.Background()) {
			var income primitive.M
			cursor.Decode(&income)

			incomes = append(incomes, domain.PrimToIncome(income).ToDTO())
		}

		logging.FoundOnDB("All Incomes", "Cloud", "model.GetIncomes")

		return incomes, nil

	}
	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.GetIncomes"))
}

func GetUserIncomes(id string) ([]domain.IncomeDTO, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.GetUserIncomes"))
	}

	if statusDBLocal {
		incomes := []domain.Income{}

		result := database.DBlocal.Where("user_id = ?", id).Find(&incomes)

		if result.Error != nil {
			logging.FailedToFindOnDB(id, "Local", result.Error, "model.GetUserIncomes")
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		}

		incomesDto := []domain.IncomeDTO{}

		for _, income := range incomes {
			incomesDto = append(incomesDto, income.ToDTO())
		}

		logging.FoundOnDB(id, "Local", "model.GetUserIncomes")
		return incomesDto, nil
	}

	if statusDBCloud {
		incomes := []domain.IncomeDTO{}

		user := bson.M{"_id": id}
		filter := bson.M{"user": user}

		cursor, err := database.DBCloud.Income.Find(context.Background(), filter)

		if err != nil {
			logging.FailedToFindOnDB(id, "Cloud", err, "model.GetUserIncomes")
			return nil, util.GetTagError(http.StatusBadRequest, err)
		}

		for cursor.Next(context.Background()) {
			var aux bson.M
			cursor.Decode(aux)
			incomes = append(incomes, domain.PrimToIncome(aux).ToDTO())
		}

		logging.FoundOnDB(id, "Cloud", "model.GetUserIncomes")
		return incomes, nil
	}
	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.GetUserIncomes"))
}

func CreateIncome(income domain.IncomeDTO) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.CreateIncome"))
	}

	if income.Value == 0 || income.Description == "" || income.ReceivedAt == "" || income.UserID == "" {
		return util.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.InvalidFields("model.CreateIncome")))
	}

	income.CreatedAt = util.GetTimeNow()
	income.UpdatedAt = util.GetTimeNow()

	if statusDBCloud {

		raw := income.ToPrim()

		result, err := database.DBCloud.Income.InsertOne(context.Background(), raw)

		if err != nil {
			logging.FailedToCreateOnDB(income.ID, "Cloud", err, "model.CreateIncome")
			return util.GetTagError(http.StatusBadRequest, err)
		}

		income.ID = result.InsertedID.(primitive.ObjectID).Hex()

		logging.CreatedOnDB(income.ID, "Cloud", "model.CreateIncome")
	}

	if statusDBLocal {

		entity := income.ToEntity()

		result := database.DBlocal.Create(&entity)

		if result.Error != nil {
			logging.FailedToCreateOnDB(income.ID, "Local", result.Error, "model.CreateIncome")
			return util.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.CreatedOnDB(income.ID, "Local", "model.CreateIncome")
		return nil
	}

	return util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.CreateIncome"))
}

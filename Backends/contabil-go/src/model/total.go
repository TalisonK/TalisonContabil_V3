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

func CreateIncomeTotal(userId string, month string, year int) (*domain.Total, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	var startingDate string
	var endingDate string

	if month == "" && year == 0 {
		startingDate, endingDate = util.GetFirstAndLastDayOfCurrentMonth()
	} else {
		startingDate, endingDate = util.GetFirstAndLastDayOfMonth(month, year)
	}

	incomes, tagError := GetIncomesByDate(userId, startingDate, endingDate)

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Incomes for user %s", userId), "Income", tagError.Inner)
		return nil, tagError
	}

	total := mountTotal(month, year, userId, incomes)

	old, tagError := findTotalByMonthAndYear(month, year)

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Totals for user %s", userId), "Total", tagError.Inner)
		return nil, tagError
	}

	if old.ID != "" {
		total.CreatedAt = old.CreatedAt
	}

	if statusDBCloud {

		if old.ID == "" {
			total, tagError = createTotal(total)
		} else {
			total, tagError = updateTotal(total, *old)
		}

		if tagError != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), "Local", tagError.Inner)
			return nil, tagError
		}

	}

	if statusDBLocal && total.ID != "" {
		result := database.DBlocal.Save(&total)

		if result.Error != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), "Local", result.Error)
			return nil, util.GetTagError(http.StatusInternalServerError, result.Error)
		}

		logging.CreatedOnDB(total.ID, "Local")
		return &total, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func createTotal(total domain.Total) (domain.Total, *util.TagError) {

	raw := total.ToPrim()

	inserted, err := database.DBCloud.Total.InsertOne(context.Background(), raw)
	if err != nil {
		logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", total.Month, total.Year), "Cloud", err)
		return domain.Total{}, util.GetTagError(http.StatusInternalServerError, err)
	}

	total.ID = inserted.InsertedID.(primitive.ObjectID).Hex()
	logging.CreatedOnDB(total.ID, "Cloud")
	return total, nil
}

func updateTotal(total domain.Total, old domain.Total) (domain.Total, *util.TagError) {

	total.ID = old.ID

	total.UpdatedAt = util.GetTimeNow()

	filter := bson.M{"_id": total.ID}

	parser := bson.M{"$set": total.ToPrim()}

	_, err := database.DBCloud.Total.UpdateOne(context.Background(), filter, parser)

	if err != nil {
		logging.FailedToUpdateOnDB(fmt.Sprintf("Total for income from %s/%d", total.Month, total.Year), "Cloud", err)
		return domain.Total{}, util.GetTagError(http.StatusInternalServerError, err)
	}

	logging.UpdatedOnDB(total.ID, "Cloud")
	return total, nil
}

func findTotalByMonthAndYear(month string, year int) (*domain.Total, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDBLocal {

		var total domain.Total

		result := database.DBlocal.Where("month = ? AND year = ?", month, year).Find(&total)

		if result.Error != nil {
			logging.FailedToFindOnDB("Income Total", "Local", result.Error)
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		}

		return &total, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())

}

func mountTotal(month string, year int, userId string, incomes []domain.Income) domain.Total {
	total := domain.Total{}

	total.CreatedAt = util.GetTimeNow()
	total.UpdatedAt = util.GetTimeNow()
	total.Type = "Income"
	total.Month = month
	total.Year = year
	total.UserID = userId

	for _, income := range incomes {
		total.TotalValue += income.Value
	}

	total.TotalValue = util.ToFixed(total.TotalValue, 2)

	return total
}

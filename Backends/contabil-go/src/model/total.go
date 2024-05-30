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
		logging.FailedToFindOnDB(userId, "Income", tagError.Inner)
		return nil, tagError
	}

	fmt.Println(incomes)

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

	if statusDBCloud {

		raw := total.ToPrim()

		old, tagErr := findTotalByMonthAndYear(month, year)

		if tagErr != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), "Local", tagErr.Inner)
			return nil, tagErr
		}

		if old.ID == "" {
			inserted, err := database.DBCloud.Total.InsertOne(context.Background(), raw)
			if err != nil {
				logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), "Cloud", err)
				return nil, util.GetTagError(http.StatusInternalServerError, err)
			}

			total.ID = inserted.InsertedID.(primitive.ObjectID).Hex()
			logging.CreatedOnDB(total.ID, "Cloud")
		} else {

			total.ID = old.ID

			filter := bson.M{"_id": total.ID}

			parser := bson.M{"$set": total.ToPrim()}

			_, err := database.DBCloud.Total.UpdateOne(context.Background(), filter, parser)

			if err != nil {
				logging.FailedToUpdateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), "Cloud", err)
				return nil, util.GetTagError(http.StatusInternalServerError, err)
			}

			logging.UpdatedOnDB(total.ID, "Cloud")

		}

	}

	if statusDBLocal && total.ID != "" {
		result := database.DBlocal.Create(&total)

		if result.Error != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), "Local", result.Error)
			return nil, util.GetTagError(http.StatusInternalServerError, result.Error)
		}

		logging.CreatedOnDB(total.ID, "Local")
		return &total, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
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

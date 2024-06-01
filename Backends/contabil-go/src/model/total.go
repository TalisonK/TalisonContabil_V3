package model

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUpdateTotal(userId string, month string, year int, totalType string) (*domain.Total, *util.TagError) {

	// check for connectivity with the databases

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	// check if the month and year are valid

	if month == "" || year == 0 {
		return nil, util.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.InvalidFields()))
	}

	// get the first and last day of the month
	startingDate, endingDate := util.GetFirstAndLastDayOfMonth(month, year)

	// fetch the incomes or expenses from the database
	var activities []domain.Activity
	var tagError *util.TagError

	if totalType == constants.INCOME {
		activities, tagError = fetchIncomesByDate(userId, startingDate, endingDate)
	} else {
		activities, tagError = fetchExpensesByDate(userId, startingDate, endingDate)
	}

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("%ss for user %s", totalType, userId), constants.INCOME, tagError.Inner)
		return nil, tagError
	}

	// mount the total
	total := mountTotal(month, year, userId, totalType, activities)

	// check if the total already exists
	old, tagError := findTotalByMonthAndYear(month, year, totalType)

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Totals for user %s", userId), "Total", tagError.Inner)
		return nil, tagError
	}

	// create or update the total in the database
	if old.ID != "" {
		total.ID = old.ID
		total.CreatedAt = old.CreatedAt
	}

	if statusDBCloud {

		if old.ID == "" {
			total, tagError = createTotalInDB(total)
		} else {
			total, tagError = updateTotalInDB(total, *old)
		}

		if tagError != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), constants.CLOUD, tagError.Inner)
			return nil, tagError
		}

	}

	if statusDBLocal {
		result := database.DBlocal.Save(&total)

		if result.Error != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for %s from %s/%d", totalType, month, year), constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusInternalServerError, result.Error)
		}

		logging.CreatedOnDB(total.ID, constants.LOCAL)
		return &total, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

// createTotalInDB creates the total in the database
func createTotalInDB(total domain.Total) (domain.Total, *util.TagError) {

	raw := total.ToPrim()

	inserted, err := database.DBCloud.Total.InsertOne(context.Background(), raw)
	if err != nil {
		logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", total.Month, total.Year), constants.CLOUD, err)
		return domain.Total{}, util.GetTagError(http.StatusInternalServerError, err)
	}

	total.ID = inserted.InsertedID.(primitive.ObjectID).Hex()
	logging.CreatedOnDB(total.ID, constants.CLOUD)
	return total, nil
}

// updateTotalInDB updates the total in the database
func updateTotalInDB(total domain.Total, old domain.Total) (domain.Total, *util.TagError) {

	total.UpdatedAt = util.GetTimeNow()

	filter := bson.M{"_id": total.ID}

	parser := bson.M{"$set": total.ToPrim()}

	_, err := database.DBCloud.Total.UpdateOne(context.Background(), filter, parser)

	if err != nil {
		logging.FailedToUpdateOnDB(fmt.Sprintf("Total for income from %s/%d", total.Month, total.Year), constants.CLOUD, err)
		return domain.Total{}, util.GetTagError(http.StatusInternalServerError, err)
	}

	logging.UpdatedOnDB(total.ID, constants.CLOUD)
	return total, nil
}

func fetchIncomesByDate(userId string, startingDate string, endingDate string) ([]domain.Activity, *util.TagError) {

	var activities []domain.Activity

	incomes, tagError := GetIncomesByDate(userId, startingDate, endingDate)

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Incomes for user %s", userId), constants.INCOME, tagError.Inner)
		return nil, tagError
	}

	for _, income := range incomes {
		activities = append(activities, income.ToActivity())
	}

	return activities, nil

}

func fetchExpensesByDate(userId string, startingDate string, endingDate string) ([]domain.Activity, *util.TagError) {

	var activities []domain.Activity

	expenses, tagError := GetExpensesByDate(userId, startingDate, endingDate)

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Expenses for user %s", userId), constants.INCOME, tagError.Inner)
		return nil, tagError
	}

	for _, expenses := range expenses {
		activities = append(activities, expenses.ToActivity())
	}

	return activities, nil

}

func findTotalByMonthAndYear(month string, year int, totalType string) (*domain.Total, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDBLocal {

		var total domain.Total

		result := database.DBlocal.Where("month = ? AND year = ? AND type = ?", month, year, totalType).Find(&total)

		if result.Error != nil {
			logging.FailedToFindOnDB("Income Total", constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		}

		return &total, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())

}

func mountTotal(month string, year int, userId string, totalType string, activities []domain.Activity) domain.Total {
	total := domain.Total{}

	total.CreatedAt = util.GetTimeNow()
	total.UpdatedAt = util.GetTimeNow()
	total.Type = totalType
	total.Month = month
	total.Year = year
	total.UserID = userId

	for _, activity := range activities {
		total.TotalValue += activity.Value
	}

	total.TotalValue = util.ToFixed(total.TotalValue, 2)

	return total
}

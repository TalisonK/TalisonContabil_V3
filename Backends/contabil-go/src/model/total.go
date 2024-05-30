package model

import (
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func CreateIncomeTotal(userId string, month string, year int) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
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
		return tagError
	}

	fmt.Println(incomes)
	return nil
}

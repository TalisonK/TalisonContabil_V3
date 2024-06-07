package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	var body domain.Total

	json.NewDecoder(r.Body).Decode(&body)

	startingDate, endingDate := util.GetFirstAndLastDayOfMonth(body.Month, body.Year)

	result, tagerr := model.GetExpensesByDate(body.UserID, startingDate, endingDate)

	if tagerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, logging.FailedToFindOnDB(fmt.Sprintf("Expenses from user %s", body.UserID), constants.LOCAL, tagerr.Inner))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

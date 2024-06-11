package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	var body domain.Total

	json.NewDecoder(r.Body).Decode(&body)

	startingDate, endingDate := timeHandler.GetFirstAndLastDayOfMonth(body.Month, body.Year)

	result, tagerr := model.GetExpensesByDate(body.UserID, startingDate, endingDate, true, true)

	if tagerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, logging.FailedToFindOnDB(fmt.Sprintf("Expenses from user %s", body.UserID), constants.LOCAL, tagerr.Inner))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func CreateExpense(w http.ResponseWriter, r *http.Request) {

	var body domain.ExpenseDTO

	json.NewDecoder(r.Body).Decode(&body)

	result, tagErr := model.CreateExpenseHandler(body)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		fmt.Fprintln(w, logging.GenericError("Error received while tring do create expense", tagErr.Inner))
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

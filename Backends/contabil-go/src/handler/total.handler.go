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

func GetTotalRange(w http.ResponseWriter, r *http.Request) {

	entry := domain.Total{}

	json.NewDecoder(r.Body).Decode(&entry)

	income, err := model.CreateUpdateTotal(entry.UserID, entry.Month, entry.Year, constants.INCOME)

	if err != nil {
		fail(w, constants.INCOME, *err, entry)
	}

	expense, err := model.CreateUpdateTotal(entry.UserID, entry.Month, entry.Year, constants.EXPENSE)

	if err != nil {
		fail(w, constants.EXPENSE, *err, entry)
	}

	results := []domain.Total{}

	results = append(results, *expense, *income)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)

}

func fail(w http.ResponseWriter, totalType string, err util.TagError, entry domain.Total) {
	message := fmt.Sprintf("Total for %s from %s/%d", totalType, entry.Month, entry.Year)
	logging.FailedToCreateOnDB(message, "Totals", err.Inner)
	w.WriteHeader(err.HtmlStatus)
	fmt.Fprint(w, message)
}

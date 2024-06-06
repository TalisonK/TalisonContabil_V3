package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func GetTotalRange(w http.ResponseWriter, r *http.Request) {

	entry := domain.Total{}

	json.NewDecoder(r.Body).Decode(&entry)

	result, err := model.TotalRanger(context.Background(), nil, entry.UserID, entry.Month, entry.Year)

	if err != nil {
		fail(w, "range", *err, entry)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func fail(w http.ResponseWriter, totalType string, err util.TagError, entry domain.Total) {
	message := fmt.Sprintf("Total for %s from %s/%d", totalType, entry.Month, entry.Year)
	logging.FailedToCreateOnDB(message, "Totals", err.Inner)
	w.WriteHeader(err.HtmlStatus)
	fmt.Fprint(w, message)
}

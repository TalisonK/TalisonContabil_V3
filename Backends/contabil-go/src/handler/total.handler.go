package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func CreateTotal(w http.ResponseWriter, r *http.Request) {

	entry := domain.Total{}

	json.NewDecoder(r.Body).Decode(&entry)

	incomes, err := model.CreateIncomeTotal(entry.ID, entry.Month, entry.Year)

	if err != nil {
		logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", entry.Month, entry.Year), "Total", err.Inner)
		return
	}

	fmt.Println(incomes)

}

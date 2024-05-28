package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
)

func GetIncomes(w http.ResponseWriter, r *http.Request) {

	result, err := model.GetFullIncomes()

	if err != nil {
		w.WriteHeader(err.HtmlStatus)
		fmt.Fprintln(w, err.Inner.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetUserIncomes(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "ID is required")
		return
	}

	result, err := model.GetUserIncomes(id)

	if err != nil {
		w.WriteHeader(err.HtmlStatus)
		fmt.Fprintln(w, err.Inner.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateIncome(w http.ResponseWriter, r *http.Request) {
	var income domain.IncomeDTO

	err := json.NewDecoder(r.Body).Decode(&income)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON")
		return
	}

	tagErr := model.CreateIncome(income)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		fmt.Fprintln(w, tagErr.Inner.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

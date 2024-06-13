package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
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

func UpdateIncome(w http.ResponseWriter, r *http.Request) {
	var income domain.IncomeDTO

	err := json.NewDecoder(r.Body).Decode(&income)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON")
		return
	}

	if income.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "ID is required")
		return
	}

	result, tagErr := model.UpdateIncome(income)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		fmt.Fprintln(w, tagErr.Inner.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	if id == "" {
		logging.GenericError("Empty id passed.", nil)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty id passed.")
		return
	}

	err := model.DeleteIncome(id)

	if err != nil {
		logging.GenericError("Failed to delete income", err.Inner)
		w.WriteHeader(err.HtmlStatus)
		fmt.Fprintln(w, "Failed to delete income")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Income deleted")
}

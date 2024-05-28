package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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

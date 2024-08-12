package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/model"
)

func GetCreditCards(w http.ResponseWriter, r *http.Request) {

	userId := r.PathValue("id")

	result := model.GetCreditCardsByUser(userId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func CreateCreditCard(w http.ResponseWriter, r *http.Request) {

	var body domain.CreditCard

	json.NewDecoder(r.Body).Decode(&body)

	result, tagErr := model.CreateCreditCard(body)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		json.NewEncoder(w).Encode(tagErr)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

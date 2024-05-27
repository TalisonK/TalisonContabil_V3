package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/model"
)

func GetIncomes(w http.ResponseWriter, r *http.Request) {

	result, err := model.GetIncomes()

	if err != nil {
		w.WriteHeader(err.HtmlStatus)
		fmt.Fprintln(w, err.Inner.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	return

}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Get categories from database

	categories, err := model.GetCategories()

	if err != nil {
		util.LogHandler("Failed to get categories", err, "handler.GetCategories")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to get categories")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
		return
	}

}

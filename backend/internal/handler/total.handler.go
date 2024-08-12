package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
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

func fail(w http.ResponseWriter, totalType string, err tagError.TagError, entry domain.Total) {
	message := fmt.Sprintf("Total for %s from %s/%d", totalType, entry.Month, entry.Year)
	logging.FailedToCreateOnDB(message, "Totals", err.Inner)
	w.WriteHeader(err.HtmlStatus)
	fmt.Fprint(w, message)
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {

	var entry domain.DashboardPacket

	json.NewDecoder(r.Body).Decode(&entry)

	out, tagErr := model.GetDashboard(entry)

	if tagErr != nil {
		logging.ErrorOccurred()
		w.WriteHeader(tagErr.HtmlStatus)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)

}

func GetActivities(w http.ResponseWriter, r *http.Request) {

	userId := r.PathValue("id")

	out, tagErr := model.GetActivities(userId)

	if tagErr != nil {
		logging.ErrorOccurred()
		w.WriteHeader(tagErr.HtmlStatus)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)

}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {

	var bucket []domain.Activity

	json.NewDecoder(r.Body).Decode(&bucket)

	result, err := model.DeleteBucket(bucket)

	if err != nil {
		logging.FailedToDeleteOnDB("bucket", "All", err.Inner)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	logging.DeletedOnDB("bucket", "all")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func UpdateBucket(w http.ResponseWriter, r *http.Request) {

	var body domain.Activity

	json.NewDecoder(r.Body).Decode(&body)

	result, err := model.UpdateBucket(body)

	if err != nil {
		logging.FailedToUpdateOnDB("bucket", "All", err.Inner)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	logging.UpdatedOnDB("bucket", "all")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func ClearCache(w http.ResponseWriter, r *http.Request) {

	model.ClearCache()

	logging.GenericSuccess("Cache cleared")
	w.WriteHeader(http.StatusOK)

}

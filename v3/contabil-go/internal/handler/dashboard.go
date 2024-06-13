package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
)

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

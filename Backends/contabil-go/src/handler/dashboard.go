package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
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

package logic

import (
	"encoding/json"
	"net/http"
)

func (l *Logic) HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

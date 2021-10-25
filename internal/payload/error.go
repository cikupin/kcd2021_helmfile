package payload

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func CreateErrorResponse(w http.ResponseWriter, err error, errMessage string) {
	errMsg := ErrorResponse{
		ErrorMessage: errMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errMsg)
}

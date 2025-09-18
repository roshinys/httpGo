package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roshinys/httpGo/internal/models"
)

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &models.ApiResponse{}
		res.Message = "Server is healthy" // ← Fix this
		res.Success = true
		w.Header().Set("Content-Type", "application/json")
		resBytes, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resBytes)
	}
}

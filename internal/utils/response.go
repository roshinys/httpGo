package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/roshinys/httpGo/internal/models"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, response *models.ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := &models.ApiResponse{
		Success: false,
		Message: message,
	}
	RespondWithJSON(w, statusCode, response)
}

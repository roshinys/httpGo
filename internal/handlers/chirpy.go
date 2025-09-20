package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roshinys/httpGo/internal/models"
	"github.com/roshinys/httpGo/internal/services"
	"github.com/roshinys/httpGo/internal/utils"
)

func ChirpyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		var chirpy models.ChirpyBody
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // Strict JSON parsing
		if err := decoder.Decode(&chirpy); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
			return
		}
		if len(chirpy.Body) > 200 {
			utils.RespondWithError(w, http.StatusBadRequest, "Body Length is too long")
			return
		}

		services.FilterBody(&chirpy)
		response := &models.ApiResponse{
			Success: true,
			Message: "Valid body",
			Data:    chirpy.Body,
		}
		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

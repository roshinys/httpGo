package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roshinys/httpGo/internal/config"
	"github.com/roshinys/httpGo/internal/models"
	"github.com/roshinys/httpGo/internal/services"
	"github.com/roshinys/httpGo/internal/utils"
)

func ChirpyHandler(cfg *config.ApiConfig) http.HandlerFunc {
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
		if len(chirpy.Email) == 0 {
			utils.RespondWithError(w, http.StatusBadRequest, "Email cannot be empty")
			return
		}
		if len(chirpy.Body) > 200 {
			utils.RespondWithError(w, http.StatusBadRequest, "Body Length is too long")
			return
		}

		newChirpy, err := services.CreateChirpy(r, cfg, &chirpy)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		response := &models.ApiResponse{
			Success: true,
			Message: "Valid body",
			Data:    newChirpy,
		}
		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

func FetchChirpHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		chirps, err := services.FetchChirps(r, cfg)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		response := &models.ApiResponse{
			Success: true,
			Message: "Fetched Chirps",
			Data:    chirps,
		}
		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

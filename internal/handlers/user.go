package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roshinys/httpGo/internal/config"
	"github.com/roshinys/httpGo/internal/models"
	"github.com/roshinys/httpGo/internal/services"
	"github.com/roshinys/httpGo/internal/utils"
)

func UserHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var userDto models.UserDto
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // Strict JSON parsing
		if err := decoder.Decode(&userDto); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
			return
		}
		if len(userDto.Email) == 0 {
			utils.RespondWithError(w, http.StatusBadRequest, "Email cannot be empty")
			return
		}

		user, err := services.CreateUser(r, &userDto, cfg)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		response := &models.ApiResponse{
			Success: true,
			Message: "Created User",
			Data:    user,
		}
		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

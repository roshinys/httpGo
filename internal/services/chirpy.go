package services

import (
	"errors"
	"net/http"

	"github.com/roshinys/httpGo/internal/config"
	"github.com/roshinys/httpGo/internal/database"
	"github.com/roshinys/httpGo/internal/models"
)

func CreateChirpy(r *http.Request, cfg *config.ApiConfig, chirpy *models.ChirpyBody) (*database.Chirpy, error) {
	userExist, err := cfg.DB.FetchUserByEmail(r.Context(), chirpy.Email)
	if err != nil {
		return nil, errors.New("user with Email doesn't exist")
	}
	params := database.CreateChirpyParams{
		Body:   chirpy.Body,
		Userid: userExist.ID,
	}
	newChirpy, err := cfg.DB.CreateChirpy(r.Context(), params)
	if err != nil {
		return nil, err
	}
	return &newChirpy, nil
}

func FetchChirps(r *http.Request, cfg *config.ApiConfig) ([]database.Chirpy, error) {
	chirps, err := cfg.DB.FetchChirps(r.Context())
	if err != nil {
		return nil, err
	}
	return chirps, nil
}

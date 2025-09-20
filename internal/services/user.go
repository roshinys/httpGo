package services

import (
	"errors"
	"net/http"

	"github.com/roshinys/httpGo/internal/config"
	"github.com/roshinys/httpGo/internal/database"
	"github.com/roshinys/httpGo/internal/models"
)

func CreateUser(r *http.Request, userDto *models.UserDto, cfg *config.ApiConfig) (*database.User, error) {
	_, err := cfg.DB.FetchUserByEmail(r.Context(), userDto.Email)
	if err == nil {
		return nil, errors.New("user with Email already exist")
	}
	user, err := cfg.DB.CreateUser(r.Context(), userDto.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

package services

import (
	"database/sql"
	"errors"
	"net/http"
	"os"

	"github.com/roshinys/httpGo/internal/auth"
	"github.com/roshinys/httpGo/internal/config"
	"github.com/roshinys/httpGo/internal/database"
	"github.com/roshinys/httpGo/internal/models"
)

func CreateUser(r *http.Request, userDto *models.UserDto, cfg *config.ApiConfig) (*database.User, error) {
	_, err := cfg.DB.FetchUserByEmail(r.Context(), userDto.Email)
	if err == nil {
		return nil, errors.New("user with Email already exist")
	}
	password, err := auth.HashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}
	params := database.CreateUserParams{
		Email: userDto.Email,
		Password: sql.NullString{
			String: password,
			Valid:  true,
		},
	}
	// Create the user
	user, err := cfg.DB.CreateUser(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func LoginUser(r *http.Request, userDto *models.UserDto, cfg *config.ApiConfig) (string, error) {
	userExist, err := cfg.DB.FetchUserByEmail(r.Context(), userDto.Email)
	if err != nil {
		return "", err
	}
	if err := auth.CheckPasswordHash(userDto.Password, userExist.Password.String); err != nil {
		return "", errors.New("password mismatch")
	}
	token, err := auth.MakeJWT(userExist.ID, os.Getenv("SECRET_KEY"), 3600)
	if err != nil {
		return "", nil
	}
	return token, nil
}

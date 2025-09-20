package services

import (
	"strings"

	"github.com/roshinys/httpGo/internal/models"
)

var badWords = [3]string{"kerfuffle", "sharbert", "fornax"}

func FilterBody(chirpy *models.ChirpyBody) {
	words := strings.Fields(chirpy.Body)
	for i, word := range words {
		for _, badWord := range badWords {
			if badWord == word {
				words[i] = "****"
			}
		}
	}
	chirpy.Body = strings.Join(words, " ")
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/roshinys/httpGo/internal/config"
)

func MetricsHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hits := cfg.GetHits()
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hits: %d", hits)))
	}
}

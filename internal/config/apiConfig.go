package config

import (
	"net/http"
	"sync/atomic"

	"github.com/roshinys/httpGo/internal/database"
)

type ApiConfig struct {
	fileserverHits atomic.Int32 // safe increments
	DB             *database.Queries
}

func (a *ApiConfig) Hit() {
	a.fileserverHits.Add(1)
}

func (a *ApiConfig) GetHits() int32 {
	return a.fileserverHits.Load()
}

func (a *ApiConfig) ResetHits() {
	a.fileserverHits.Store(0)
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.Hit()
		next.ServeHTTP(w, r)
	})
}

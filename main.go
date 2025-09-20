package main

import (
	"log"
	"net/http"
	"time"

	"github.com/roshinys/httpGo/internal/config"
	"github.com/roshinys/httpGo/internal/handlers"
)

func main() {
	cfg := &config.ApiConfig{}
	mux := http.NewServeMux()

	mux.Handle("/app", http.FileServer(http.Dir("./")))
	// mux.Handle("/assets", http.FileServer(http.Dir("./assets")))
	mux.Handle("POST /api/health", cfg.MiddlewareMetricsInc(handlers.HealthHandler()))
	mux.Handle("GET /api/metrics", handlers.MetricsHandler(cfg))
	mux.Handle("POST /api/reset", cfg.MiddlewareMetricsInc(handlers.ResetHandler(cfg)))
	mux.Handle("POST /api/validate-chirp", cfg.MiddlewareMetricsInc(handlers.ChirpyHandler()))

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server starting on http://localhost:8080")
	log.Println("Visit http://localhost:8080 to see the 404 (no handlers registered yet)")
	log.Fatal(server.ListenAndServe())
}

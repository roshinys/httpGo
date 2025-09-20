package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/roshinys/httpGo/internal/config"
	"github.com/roshinys/httpGo/internal/database"
	"github.com/roshinys/httpGo/internal/handlers"

	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")
	fmt.Printf(" DB URL is %s", dbURL)
	fmt.Println()

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(" Failed to Connect to db url ")
		return
	}
	defer db.Close()
	fmt.Println(db)
	dbQueries := database.New(db)
	cfg := &config.ApiConfig{
		DB: dbQueries,
	}
	mux := http.NewServeMux()

	mux.Handle("/app", http.FileServer(http.Dir("./")))
	// mux.Handle("/assets", http.FileServer(http.Dir("./assets")))
	mux.Handle("POST /api/health", cfg.MiddlewareMetricsInc(handlers.HealthHandler()))
	mux.Handle("GET /api/metrics", handlers.MetricsHandler(cfg))
	mux.Handle("POST /api/reset", cfg.MiddlewareMetricsInc(handlers.ResetHandler(cfg)))

	// Auth Routes
	mux.Handle("POST /api/login", cfg.MiddlewareMetricsInc(handlers.AuthLogin(cfg)))

	// User Routes
	mux.Handle("POST /api/users", cfg.MiddlewareMetricsInc(handlers.UserHandler(cfg)))

	// Chirp Routes
	mux.Handle("POST /api/chirps", cfg.MiddlewareMetricsInc(handlers.ChirpyHandler(cfg)))
	mux.Handle("GET /api/chirps", cfg.MiddlewareMetricsInc(handlers.FetchChirpHandler(cfg)))

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server starting on http://localhost:" + port)
	log.Printf("Visit http://localhost:%s to see the 404 (no handlers registered yet)", port)
	log.Fatal(server.ListenAndServe())
}

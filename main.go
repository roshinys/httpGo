package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type MyHandler struct{}

type apiConfig struct {
	fileserverHits atomic.Int32 // safe increments
}

func (a *apiConfig) hit() {
	a.fileserverHits.Add(1)
}

func (a *apiConfig) getHits() int32 {
	return a.fileserverHits.Load()
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.hit()
		next.ServeHTTP(w, r)
	})
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := &ApiResponse{}
	res.Message = "Server is Health"
	res.Success = true
	w.Header().Set("Content-Type", "application/json")
	resBytes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resBytes)
}

func main() {
	cfg := &apiConfig{}
	mux := http.NewServeMux()

	// mux.Handle("/", http.FileServer(http.Dir("./")))
	// mux.Handle("/assets", http.FileServer(http.Dir("./assets")))

	mux.Handle("/health", cfg.middlewareMetricsInc(&MyHandler{}))

	mux.Handle("/stats", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits := cfg.getHits()
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Total hits: %d", hits)))
	}))

	// mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("OK"))
	// })

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

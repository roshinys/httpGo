package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/roshinys/httpGo/internal/server"
)

func main() {
	port := 42069
	server, err := server.Serve(port)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	log.Println("Server started on port", port)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")

}

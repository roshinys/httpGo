package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/roshinys/httpGo/internal/request"
	"github.com/roshinys/httpGo/internal/response"
	"github.com/roshinys/httpGo/internal/server"
)

func main() {
	port := 42069

	handler := func(w io.Writer, req *request.Request) *response.HandlerError {
		switch req.RequestLine.RequestTarget {
		case "/yourproblem":
			body := `<html>
<head>
<title>400 Bad Request</title>
</head>
<body>
<h1>Bad Request</h1>
<p>Your request honestly kinda sucked.</p>
</body>
</html>`
			return server.NewHandlerError(response.BadRequest, body)
		case "/myproblem":
			body := `<html>
  <head>
    <title>500 Internal Server Error</title>
  </head>
  <body>
    <h1>Internal Server Error</h1>
    <p>Okay, you know what? This one is on me.</p>
  </body>
</html>`
			return server.NewHandlerError(response.InternalServer, body)

		default:
			body := `
			<html>
  <head>
    <title>200 OK</title>
  </head>
  <body>
    <h1>Success!</h1>
    <p>Your request was an absolute banger.</p>
  </body>
</html>`
			w.Write([]byte(body))
			return nil
		}
	}

	server, err := server.Serve(port, handler)
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

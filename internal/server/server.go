package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync/atomic"

	"github.com/roshinys/httpGo/internal/request"
)

type Server struct {
	port     int
	listener net.Listener
	closed   atomic.Bool
}

func Serve(port int) (*Server, error) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	s := &Server{
		listener: listener,
		port:     port,
	}
	go s.listen()
	return s, err
}

func (s *Server) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		return s.listener.Close()
	}
	return nil
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if s.closed.Load() {
				// Server closed → stop loop
				return
			}
			log.Printf("Accept error: %v", err)
			continue
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	fmt.Println("Accepted connection from", conn.RemoteAddr())
	req, err := request.RequestFromReader(conn)
	if err != nil {
		log.Printf("Failed to parse request: %v", err)
		return
	}
	reqLine := req.RequestLine
	fmt.Printf("Request Line : Version : %s , Method : %s , Target : %s \n", reqLine.HttpVersion, reqLine.Method, reqLine.RequestTarget)
	fmt.Println("Headers:")
	for k, v := range req.Headers {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Printf("Body : %s \n", string(req.Body))

	// Write a basic HTTP response
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 12\r\n" +
		"\r\n" +
		"Hello World!"

	conn.Write([]byte(response))

	fmt.Println("Connection closed")
}

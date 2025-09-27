package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync/atomic"

	"github.com/roshinys/httpGo/internal/request"
	"github.com/roshinys/httpGo/internal/response"
)

type Server struct {
	port     int
	listener net.Listener
	closed   atomic.Bool
}

type Handler func(w io.Writer, req *request.Request) *response.HandlerError

func NewHandlerError(statusCode response.StatusCode, message string) *response.HandlerError {
	return &response.HandlerError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func Serve(port int, handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	s := &Server{
		listener: listener,
		port:     port,
	}
	go s.listen(handler)
	return s, err
}

func (s *Server) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		return s.listener.Close()
	}
	return nil
}

func (s *Server) listen(handler Handler) {
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
		go s.handle(conn, handler)
	}
}

func (s *Server) handle(conn net.Conn, handler Handler) {
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

	buf := &bytes.Buffer{}
	responseWriter := response.NewWrite(conn)

	if handlerErr := handler(buf, req); handlerErr != nil {
		responseWriter.WriteHandlerError(*handlerErr)
		return
	}

	body := buf.Bytes()
	headers := response.GetDefaultHeaders(len(body))
	responseWriter.WriteStatusLine(response.StatusOk)
	responseWriter.WriteHeaders(headers)
	responseWriter.WriteBody(body)

	fmt.Println("Connection closed")
}

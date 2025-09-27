package main

import (
	"fmt"
	"log"
	"net"

	"github.com/roshinys/httpGo/internal/request"
)

func main() {
	// Create a buffer of length 8 bytes
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		req, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal(err)
		}
		reqLine := req.RequestLine
		fmt.Printf("Request Line : Version : %s , Method : %s , Target : %s \n", reqLine.HttpVersion, reqLine.Method, reqLine.RequestTarget)
		fmt.Println("Headers:")
		for k, v := range req.Headers {
			fmt.Println("Key:", k, "Value:", v)
		}
		fmt.Println("Connection closed")
	}

}

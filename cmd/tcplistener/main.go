package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	buffer := make([]byte, 8)
	ch := make(chan string, 8)
	s := ""
	go func() {
		defer f.Close()
		defer close(ch)
		for {
			n, err := f.Read(buffer)
			if err != nil {
				if s != "" {
					ch <- s
					// fmt.Printf("Read %s\n", s)
				}
				if err == io.EOF {
					// fmt.Println("Reached end of file")
					break
				}
				log.Fatal(err)
			}

			for i := 0; i < n; i++ {
				if buffer[i] == '\n' {
					// fmt.Printf("Read %s\n", s)
					ch <- s
					s = ""
				} else {
					s += string(buffer[i])
				}
			}
		}
	}()
	return ch
}

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
		ch := getLinesChannel(conn)
		for line := range ch {
			fmt.Printf("read %s\n", line)
		}
		fmt.Println("Connection closed")
	}

}

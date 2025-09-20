package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	buffer := make([]byte, 8)
	ch := make(chan string, 8)
	defer close(ch)
	s := ""
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
	return ch
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffer of length 8 bytes
	ch := getLinesChannel(file)
	for line := range ch {
		fmt.Printf("read %s\n", line)
	}
	// for {
	// 	line, ok := <-ch
	// 	if !ok {
	// 		fmt.Println("Channel closed")
	// 		break
	// 	}
	// 	fmt.Println(line)
	// }
}

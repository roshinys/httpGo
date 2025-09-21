package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAdr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := net.DialUDP("udp", nil, udpAdr)
	defer conn.Close()
	fmt.Println("udp setup connection established")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(">")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("No more user input")
			break
		}

		if message == "" {
			continue
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Failed to write to server")
			log.Fatal(err)
			break
		}

	}
}

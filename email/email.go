package email

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Println("Message Received %s", message)
	}()
	fmt.Println("Message Sent %s", message)
}

func Test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("==================>")
}

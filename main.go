package main

import "github.com/roshinys/httpGo/blockthreadex"

func main() {
	// email.Test("Hello")
	// email.Test("good")
	// email.Test("hehe")

	// emails := [3]blockthreadex.Email{
	// 	{Body: "First email content", Date: time.Now()},
	// 	{Body: "Second email content", Date: time.Now().Add(-24 * time.Hour)},
	// 	{Body: "Third email content", Date: time.Now().Add(-48 * time.Hour)},
	// }
	// res := blockthreadex.CheckEmailAge(emails)
	// fmt.Println(res)
	// ch := blockthreadex.TestAddEmails()
	// for i := 0; i < 3; i++ {
	// 	res := <-ch
	// 	fmt.Println(res)
	// }

	// blockthreadex.TestCloseEx()
	// blockthreadex.TestConFib(100)

	blockthreadex.TestSelect()

}

package main

import (
	"errors"
	"fmt"
	"strings"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cus *customer, ts transaction) (float64, error) {
	switch ts.transactionType {
	case transactionDeposit:
		cus.balance += ts.amount
	case transactionWithdrawal:
		cus.balance -= ts.amount
	default:
		return 0, errors.New("Unknow Tranastion Type")
	}
	return cus.balance, nil
}

func main() {
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

	updateBalance(&alice, deposit)

	fmt.Println(alice)

	// e := email{
	// 	message:     "Hey",
	// 	fromAddress: "from",
	// 	toAddress:   "to",
	// }

	// fmt.Println(e.message)
	// e.setMessage("Hi")
	// fmt.Println(e.message)

	// var s *string
	// if s != nil {
	// 	fmt.Println(*s)
	// }

	// var x int = 50
	// var y *int = &x
	// *y = 100

	// fmt.Println(x) // Output to be 100

	//initialize pointers
	// var p *int
	// fmt.Printf(" Lets check this out %v", p)
	// fmt.Println()

	// // lets play around

	// s := "hey hello"

	// sa := &s

	// fmt.Printf("Address of variable s is %v and its value is %v", sa, *sa)

	// fmt.Println()

	// message := "Hey how are you ?"
	// removeProfanity(&message)

	// x := 11
	// incr(x)
	// fmt.Println("Outer Pass by Value x  = %d", x)

	// y := &x
	// incrPointer(y)
	// fmt.Println("Inner Increament Pointer %d", *y)

}

func incrPointer(y *int) {
	*y++
	fmt.Println("Outer Increament Pointer %d", *y)
}

func incr(x int) {
	x++
	fmt.Println("Inner Pass by Value x  = %d", x)

}

func removeProfanity(message *string) {
	fmt.Println(*message)
	fmt.Println(&message)
	fmt.Println(len(*message))

	// ReAssign
	sa := strings.Split(*message, " ")
	ln := len(sa)
	for i := 0; i < ln; i++ {
		if sa[i] == "Hey" {
			sa[i] = "***"
		}
	}
	*message = strings.Join(sa, " ")

	fmt.Println(*message)
	fmt.Println(&message)

}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

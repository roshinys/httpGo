package main

//slice test

// func getMessageCosts(messages []string) []float64 {
// 	msgLen := len(messages)
// 	result := make([]float64, msgLen)
// 	for i := 0; i < msgLen; i++ {
// 		result[i] = float64(len(messages[i])) * 0.01
// 	}
// 	return result
// }

// func isPrime(n int) bool {
// 	res := false

// 	for i := 2; i < n; i++ {
// 		fmt.Println(n, i, n/i)
// 		if n%i == 0 {
// 			res = true
// 			break
// 		}
// 	}

// 	return res
// }

// func main() {

// exdelete()

// fmt.Println(getUserMap([]string{"roshin", "charlie"}, []int{1, 2}))

// Slice Testing
// fmt.Println(getMessageCosts([]string{"Hello", "World", "!"}))
// message, err := getMessageWithRetriesForPlan("free", [3]string{"Hello", "World", "!"})
// if err != nil {
// 	fmt.Println("Error:", err)
// }
// fmt.Println(message)

// //Arrays Testing
// pl, sl := getMessageRetries("Hello", "World", "!")
// fmt.Println(pl)
// fmt.Println(sl)

// x := 10
// res := getConnections(x)

// fmt.Printf("Connections for %d is %d\n", x, res)

// x := 10
// // Test 1: Normal division
// result, err := divide(float64(x), 2.0)
// if err != nil {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Printf("%.0f / 2 = %.2f\n", float64(x), result)
// }

// // Test 2: Division by zero (will trigger custom error)
// result2, err2 := divide(float64(x), 0.0)
// if err2 != nil {
// 	fmt.Println("Error:", err2)
// } else {
// 	fmt.Printf("%.0f / 0 = %.2f\n", float64(x), result2)
// }
// }

package blockthreadex

import "fmt"

func concurrentFib(n int) []int {
	res := []int{}
	s := make(chan int, n)
	go fibonacci(n, s)

	for item := range s {
		res = append(res, item)
	}
	return res

}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func TestConFib(n int) {
	fmt.Println(concurrentFib(n))
}

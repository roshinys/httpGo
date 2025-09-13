package blockthreadex

import "fmt"

func countReports(numSentCh chan int) int {
	go sendReports(100, numSentCh)
	res := 0
	for {
		v, ok := <-numSentCh
		if !ok {
			break
		}
		res += v
	}
	return res
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func TestCloseEx() {
	ch := make(chan int)
	fmt.Println(countReports(ch))
}

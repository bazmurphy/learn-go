package main

import (
	"fmt"
	"time"
)

// the range keyword works for channels the same as for slices and maps
// so we can range over the channel and it will block until a value is ready
// it will read it into the variable item and then it will execute the body

// for item := range ch {
//     // item is the next value received from the channel
// }

// and it will do that over and over for each value coming across the channel
// and will only exit the loop once the channel is closed

func concurrentFib(n int) {
	chInts := make(chan int)
	go func() {
		// the fibonacci function will send its results over the channel
		fibonacci(n, chInts)
	}()
	// we loop over the channel using the range keyword
	// the range loop automatically detects when the channel is closed, and it terminates the loop accordingly
	for item := range chInts {
		fmt.Println(item)
	}
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

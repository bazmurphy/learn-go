package main

import (
	"fmt"
	"time"
)

// Fix the bug in the pingPong function.
// It shouldn't return until the entire game of ping pong is complete.
// Look at the print statements to get an idea of how pinger and ponger should interact.

func pingPong(numPings int) {
	// create two channels of type empty struct (signal)
	pings := make(chan struct{})
	pongs := make(chan struct{})
	// create a (concurrent) goroutine with the function ponger
	// passing in the pings and pongs channels
	go ponger(pings, pongs)
	// call the function pinger (in the main goroutine)
	// passing in the pings and pongs channels, and number of pings
	pinger(pings, pongs, numPings)
}

// don't touch below this line

func pinger(pings, pongs chan struct{}, numPings int) {

	fmt.Println("PINGER RAN")

	// create a goroutine
	go func() {
		fmt.Println("PINGER GO FUNC RAN")

		// sleep for 50ms
		sleepTime := 50 * time.Millisecond
		// loop over the number of pings
		for i := 0; i < numPings; i++ {
			fmt.Printf("sending ping %v\n", i)
			// send an empty struct "signal" to the pings channel
			pings <- struct{}{}
			// sleep 50 ms
			time.Sleep(sleepTime)
			// sleep 100ms (50 * 2)
			sleepTime *= 2
		}
		// close the pings channel
		close(pings)
	}() // IIFE
	i := 0
	for range pongs { // loop over the pongs channel and receive the value
		fmt.Println("got pong", i)
		// increment i
		i++
	}
	fmt.Println("pongs done")
}

func ponger(pings, pongs chan struct{}) {

	fmt.Println("PONGER RAN")

	i := 0
	for range pings { // loop over the pings channel and receive the value
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		// send an empty struct "signal" to the pongs channel
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	// close the pongs channel
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}

package main

import (
	"fmt"
	"time"
)

// a token is a unary value, it is not binary (two possible values, true/false), there is just one possible value
// and when there is just one possible value we don't care much about that value, its not interesting in any way
// we don't get what is passed, we just care when and if something is passed through the channel

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		// receive the token
		<-dbChan
	}
}

// don't touch below this line

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	// create a new goroutine
	go func() {
		for i := 0; i < numDBs; i++ {
			// send the token (an empty struct)
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}() // < note IIFE
	return ch
}

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	test(3)
	test(4)
	test(5)
}

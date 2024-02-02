package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	// use the go keyword to make the function concurrent
	// the go keyword immediately moves down to mt.Printf("Email sent~) on the main thread, the main go routine
	// and spawn a new go routine to do the go func()
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Stacy!")
	test("Hi there John!")
	test("Hey there Jane!")
}

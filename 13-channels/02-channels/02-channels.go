package main

import (
	"fmt"
	"time"
)

// (!) sending and receiving values from a channel are BLOCKING operations
// if we are trying to send a value into a channel and there is no other goroutine on the other side to read/receive it out
// then the code will stop and wait on the goroutine until there is a reader ready
// and the same goes for receiving, if there is nothing being sent into the channel on another go routine
// then the code will stop and wait until there is something sent

// a deadlock is when all of the goroutines in a program are blocking and there is nothing for them to do

func filterOldEmails(emails []email) {
	// we create a channel
	isOldChan := make(chan bool)

	// we create a new goroutine
	// that loops over the emails and sends boolean values into the channel
	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				// send the value to the channel
				isOldChan <- true
				continue
			}
			// send the value to the channel
			isOldChan <- false
		}
	}() // < note this is an IIFE

	// and then the code below will execute at the same time

	// receive the value from the channel
	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	// receive the value from the channel
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	// receive the value from the channel
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}

// TEST SUITE -- Don't touch below this line
type email struct {
	body string
	date time.Time
}

func test(emails []email) {
	filterOldEmails(emails)
	fmt.Println("==========================================")
}

func main() {
	test([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Yo are you okay?",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Have you heard of that website Boot.dev?",
			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "It's awesome honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Today is the day!",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What do you want for lunch?",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Why are you the way that you are?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Did we do it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Letsa Go!",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Okay...?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
}

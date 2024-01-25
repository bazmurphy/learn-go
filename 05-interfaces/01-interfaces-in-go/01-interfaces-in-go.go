package main

import (
	"fmt"
	"time"
)

// define a function takes in a msg (which is an interface)
func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

// define an interface with a single method
type message interface {
	getMessage() string
}

// don't edit below this line

// define a struct with two fields
type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

// implement the getMessage method for birthdayMessage
func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

// define a struct with two fields
type sendingReport struct {
	reportName    string
	numberOfSends int
}

// implement the getMessage method for sendingReport
func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// the test function is not passed "interface literals" (which aren't even a thing)
// (an interface is an abstract type that represents other types)
// because the test function takes an interface
// we can pass it any struct that implements that interface
// so in this case we can pass it sendingReport or birthdayMessage because they both implement that interface
// those are two different types in a strongly typed language
// both being passed in as the first parameter to a single function
// the reason it works is because we are using interfaces
// (!) the test function takes a message interface and demonstrates polymorphism
func test(m message) {
	sendMessage(m)
	fmt.Println("====================================")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}

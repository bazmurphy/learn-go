package main

import "fmt"

// in the previous channels examples - there is never anything stored in the channel
// we send, and there must be a receiver
// but we can allow channels to store values
// so senders can send into the channel even when there is no receivers
// when a receiver goes to read a value from the channel, it will read them in order one by one (as they were sent in)

// Channels can optionally be buffered
// You can provide a buffer length as the second argument to make() to create a buffered channel:
// ch := make(chan int, 100)
// Sending on a buffered channel only blocks when the buffer is full.
// Receiving blocks only when the buffer is empty.

// in this way we can do "batching"

func addEmailsToQueue(emails []string) chan string {
	// we specify a buffer length (a batch size)
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	// note that they (add and send) are both running in the same goroutine
	// add emails is called
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	// send email is called
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func main() {
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}

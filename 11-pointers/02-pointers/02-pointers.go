package main

import (
	"fmt"
	"strings"
)

// the * is used in a pointers type eg:
// var myString *string

// myString := "hello"

// the & is used to generate a pointer
// myStringPointer = &myString

// the * is also used in a (!)different way to dereference a pointer to gain access to the underlying value

// read myString through the pointer
// fmt.Println(*myStringPointer)

// set myString through the pointer
// *myStringPointer = "world"

// the parameter is a pointer
func removeProfanity(message *string) {
	// messageVal is a copy of the value pointed to by message
	messageVal := *message
	// we make changes to messageVal (but this doesn't update the original message value)
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	// set the message through the pointer
	*message = messageVal
}

// func removeProfanity(message *string) {
// 	profanities := []string{"dang", "shoot", "heck"}
// 	for _, profanity := range profanities {
// 		*message = strings.ReplaceAll(*message, profanity, strings.Repeat("*", len(profanity)))
// 	}
// }

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}

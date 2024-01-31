package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

// two common reasons to use a pointer in go
// 1. you want to pass a value into a function and change the value and have those changes persist outside of the function
// (normally when you pass a value into a function a copy is passed in)
// 2. performance optimisations:
// every time you create a copy of a variable in memory you have to copy that variable in memory, which takes time and space
// if you are using lots of data you can make micro-optimisations and use pointers if you want to avoid all the memory copying
// pointers are dangerous in that they can lead to bugs if not used properly

// Don't touch above this line

func sendMessage(m Message) {
	// fmt.Printf("To: %v\n", &m.Recipient)
	fmt.Printf("To: %v\n", m.Recipient)
	// fmt.Printf("Message: %v\n", &m.Text)
	fmt.Printf("Message: %v\n", m.Text)
}

// Don't touch below this line

func test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=====================================")
}

func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")
}

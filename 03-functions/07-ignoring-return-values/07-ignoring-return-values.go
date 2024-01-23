package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
	// Welcome to Textio, John
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}

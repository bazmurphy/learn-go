package main

import "fmt"

func main() {
	// we create variable x as a type integer with the value of 50
	var x int = 50
	// we create y which is a pointer of type integer which references x
	var y *int = &x
	// we deference y and set it to 100
	// so we set x's value through the pointer y
	*y = 100

	fmt.Println("x is: ", x)
	// x is: 100
	fmt.Println("*y is: ", *y)
	// y is: 100
	fmt.Println("y is: ", y)
	// y is: 0xc0000120f8
}

// What is the value of *y after the code executes?

// a memory address pointing to x

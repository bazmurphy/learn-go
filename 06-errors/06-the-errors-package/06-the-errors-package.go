package main

import (
	"errors" // this is the errors package in the go standard library
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		// we can create a new error from a string
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

// don't edit below this line

func test(x, y float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", x, y)
	quotient, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}

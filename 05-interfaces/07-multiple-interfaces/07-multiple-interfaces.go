package main

import (
	"fmt"
)

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

// My (More Verbose) Version:
// func (e email) cost() float64 {
// 	// get the length of the email
// 	characters := len(e.body)

// 	// initialise the cost variable
// 	var cost float64
// 	// establish the cost
// 	if e.isSubscribed {
// 		cost = 0.01
// 	} else {
// 		cost = 0.05
// 	}

// 	// calculate the totalCost
// 	totalCost := float64(characters) * cost

// 	return totalCost
// }

func (e email) print() {
	fmt.Println(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

// the test function takes an expense interface and a printer interface
// which have the cost method and print method on them respectively
// we are passing instances of the email struct into test()
// as BOTH the expense and the printer because the email struct implements BOTH interfaces
func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}

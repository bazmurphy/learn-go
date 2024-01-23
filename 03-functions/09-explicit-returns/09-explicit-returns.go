package main

import (
	"fmt"
)

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

// don't edit below this line

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================")
}

func main() {
	test(4)
	// Age: 4
	// You are an adult in 14 years
	// You can drink in 17 years
	// You can rent a car in 21 years
	// ====================================
	test(10)
	// Age: 10
	// You are an adult in 8 years
	// You can drink in 11 years
	// You can rent a car in 15 years
	// ====================================
	test(22)
	// Age: 22
	// You are an adult in 0 years
	// You can drink in 0 years
	// You can rent a car in 3 years
	// ====================================
	test(35)
	// Age: 35
	// You are an adult in 0 years
	// You can drink in 0 years
	// You can rent a car in 0 years
	// ====================================
}

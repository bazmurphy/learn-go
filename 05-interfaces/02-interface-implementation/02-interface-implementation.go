package main

import (
	"fmt"
)

// Interfaces are implemented implicitly
// As long as a struct satisfies all the requirements of an interface
// then it that struct implements that interface automatically

// define an interface with two methods
type employee interface {
	getName() string
	getSalary() int
}

// define a struct with three fields
type contractor struct {
	name string
	// in contract the salary is a calculation of two variables
	hourlyPay    int
	hoursPerYear int
}

// implement the getName method on the contractor struct
func (c contractor) getName() string {
	return c.name
}

// implement the getSalary method on the contractor struct
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

// don't touch below this line

type fullTime struct {
	name string
	// in fullTime the salary is stored as an int
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

func main() {
	test(fullTime{
		name:   "Jack",
		salary: 50000,
	})
	test(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})
}

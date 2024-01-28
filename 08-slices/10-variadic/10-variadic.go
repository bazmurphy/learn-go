package main

import "fmt"

// ...float64 as far as the function definition is concerned is just a slice
// we treat this the same as if it were []float64
// but you can pass in any number of arguments and they will come in as a slice
// it gives more flexibility
func sum(nums ...float64) float64 {
	var total float64
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

// the spread operator we take a slice of values and pass them into a variadic function
// (nums...)

// don't edit below this line

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}

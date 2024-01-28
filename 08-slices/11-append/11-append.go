package main

import "fmt"

type cost struct {
	day   int
	value float64
}

// func getCostsByDay(costs []cost) []float64 {
// 	output := make([]float64, costs[len(costs)-1].day+1)
// 	for i := 0; i < len(costs); i++ {
// 		entry := costs[i]
// 		output[entry.day] += entry.value
// 	}
// 	return output
// }

// append(slice, oneThing)
// append(slice, firstThing, secondThing)
// append(slice, anotherSlice...)

func getCostsByDay(costs []cost) []float64 {
	// slice literal syntax
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		// while the day is greater than the length of costsByDay
		// if we encounter a day that we don't have room from
		// then we should grow the slice by appending 0.0 until we have room
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		// if you try to index into an index that is outside the length of the slice
		// eg. if you have a slice of length 3 and you try to access index 6
		// then you will get this error: "panic: runtime error: index out of range [0] with length 0"
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}

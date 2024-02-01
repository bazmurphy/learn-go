package main

import (
	"fmt"
	"time"

	// we use "go get"
	// "go.sum" contains dependencies
	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}

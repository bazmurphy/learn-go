package main

import (
	"fmt"
	"time"
)

// the select statement also has a default case
// you don't need to always use it
// (!) the default case only fires if you are interested in non-blocking

// select {
//   case v := <-ch:
//     // use v
//   default:
//     // receiving from ch would block
//     // so do something else
// }

// the v value is puled from the channel IF there is a value to be pulled at the time we enter the select block
// if there is not a value ready to be read out of the channel, then the default case is fired immediately
// so default turns the select block into a non blocking block of code

// time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
// time.After() sends a value once after the duration has passed.
// time.Sleep() blocks the current goroutine for the specified amount of time.

// we can take a channel and pass it into a function
// but specify within the function that we are only going to read from the channel
// within the function we will not be able to write to the channel

// func readCh(ch <-chan int) {
// 	// ch can only be read from in this function
// }

// func writeCh(ch chan<- int) {
// 	// ch can only be written to in this function
// }

// readers read OUT OF the channel so the syntax is : func readCh(ch <-chan int) {
// writers write IN TO the channel so the syntax is : func writeCh(ch chan<- int) {

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		// When the snapshotTicker channel sends a signal (every 800 milliseconds),
		// take a snapshot by calling the takeSnapshot() function
		case <-snapshotTicker:
			takeSnapshot()
		// When the saveAfter channel sends a signal (after 2800 milliseconds),
		// save the snapshot by calling the saveSnapshot() function then return
		case <-saveAfter:
			saveSnapshot()
			return
		// If neither snapshotTicker nor saveAfter has sent a signal,
		// there is no immediate work to do, Wait for data and sleep for 500 milliseconds
		default:
			waitForData()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// don't touch below this line

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func main() {
	test()
}

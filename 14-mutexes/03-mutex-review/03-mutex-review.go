// How many threads can Lock a mutex at once?

// 1

// This example creates a map, then starts two goroutines which each have access to the map.
// One goroutine continuously mutates the values stored in the map, while the other prints the values it finds in the map.
// If we run the program on a multi-core machine, we get the following output: fatal error: concurrent map iteration and map write
// In Go, it isn’t safe to read from and write to a map at the same time.
// Mutexes to the rescue

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	m := map[int]int{}
// 	go writeLoop(m)
// 	go readLoop(m)

// 	// stop program from exiting, must be killed
// 	block := make(chan struct{})
// 	<-block
// }

// func writeLoop(m map[int]int) {
// 	for {
// 		for i := 0; i < 100; i++ {
// 			m[i] = i
// 		}
// 	}
// }

// func readLoop(m map[int]int) {
// 	for {
// 		for k, v := range m {
// 			fmt.Println(k, "-", v)
// 		}
// 	}
// }

// In the next example, we added a sync.Mutex{} and named it mux.
// In the write loop, the Lock() method is called before writing, and then the Unlock() is called when we're done.
// This Lock/Unlock sequence ensures that no other threads can Lock() the mutex while we have it locked
// – any other threads attempting to Lock() will block and wait until we Unlock().

// In the reader, we Lock() before iterating over the map,
// and likewise Unlock() when we're done.
// Now the threads share the memory safely!

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mux := &sync.Mutex{}

	go writeLoop(m, mux)
	go readLoop(m, mux)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock()
			m[i] = i
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.Mutex) {
	for {
		mux.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.Unlock()
	}
}

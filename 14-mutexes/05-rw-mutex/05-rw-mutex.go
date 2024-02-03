package main

import (
	"fmt"
	"sort"
	"sync"
)

// there is also a Sync.RWMutex

// it has the same Lock() and Unlock() as Sync.Mutex
// RLock() and RUnlock()

// The reason you would use a Read Write Mutex
// because it allows you to have multiple readers at the same time

// It is actually safe to read from a map with multiple goroutines at the same time
// as long as you are not reading AND writing at the same time
// or writing and writing at the same time
// because writing is the dangerous operation

// if the Mutex is Lock() then no other goroutines can Lock() or RLock() the Mutex
// if the Mutex is just RLock() then other goroutines can RLock() it
// So you can get multiple readers, but you can't get readers and writers, or writers and writers
// It allows many goroutines to concurrently read
// If you have a program with a shared resource,
// where most of the threads are interested with reading the shared resource,
// and you only have a few writers
// You can make your code a lot faster by using a RWMutex and allowing multiple readers at the same time
// Your code will move faster and have fewer Locks slowing down your application

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	// RLock the mutex, it can now be read by multiple goroutines
	sc.mu.RLock()
	// RUnlock the mutex
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	// time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest struct {
	email string
	count int
}

func test(sc safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	sc.mu.RLock()
	defer sc.mu.RUnlock()
	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mu:     &sync.RWMutex{},
	}
	test(sc, []emailTest{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	test(sc, []emailTest{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}

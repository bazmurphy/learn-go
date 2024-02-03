package main

import (
	"fmt"
	"sort"
	"sync"
)

// Mutexes are a primitive built into the standard library
// they allow us to communicate/share data between two go routines
// Mutexes work by locking access to protected resources
// so that only one goroutine can access that resource at a time
// sync.Mutex type is exposed by the standard library
// you create a new Mutex and you share it across many different goroutines
// then you wrap whatever code/resource that is dangerous
// that you never want two (or more) goroutines getting access to at the same time
// you wrap them in a call to .Lock() and .Unlock()

// so we can use a function to lock the mutex for a specific go routine
// and when the function finishes we unlock the mutex
// every other goroutine that calls that function will block at the .Lock() function call
// because the Mutex is blocked by another goroutine
// Mutex stands for Mutual Exclusion because it excludes every goroutine except one (the one that holds the Lock)
// and once that goroutine Unlocks the Mutex, another goroutine is able to pickup the Lock, execute its code and move on from there

// so it's a way to take some dangerous resource and share it across many different goroutines safely

// there are a couple of reasons why would you want to protect a resource and only allow a single goroutine to work on that resource at a time
// one common one is maps, maps are not thread safe in go
// if two different go routines are trying to read from and write to the same map, go will PANIC
// because that has a high likelihood of causing a race condition (when two different goroutines racing to get access to a specific resource)

// imagine we have two goroutines both trying to access a count variable and perform some operation on it

// for example:
// count = 5
// countDoubled = count * count
// count = countDoubled

// if both goroutines access this variable at the same time and perform this operation
// count will still be 25
// when maybe we wanted this operation to happen twice (because we specifically used two goroutines)
// 5 * 5, then 25 * 25 = 625

// So we could wrap that count variable in a Mutex to avoid it being accessed at the same time by different goroutines

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

// don't touch above this line

// locks the counts map
// increments it
// then unlocks it
func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

// locks the counts map
// reads from it
// then unlocks the counts map
func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	// time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	// time.Sleep(time.Microsecond)
	return sc.counts[key]
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

	for _, email := range emailsSorted {
		go sc.inc(email)
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mu:     &sync.Mutex{},
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

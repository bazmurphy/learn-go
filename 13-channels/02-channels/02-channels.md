# Channels

Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.

## Create a channel

Like maps and slices, channels must be created _before_ use. They also use the same `make` keyword:

```go
ch := make(chan int)
```

## Send data to a channel

```go
ch <- 69
```

The `<-` operator is called the _channel operator_. Data flows in the direction of the arrow. This operation will _block_ until another goroutine is ready to receive the value.

## Receive data from a channel

```go
v := <-ch
```

This reads and removes a value from the channel and saves it into the variable `v`. This operation will _block_ until there is a value in the channel to be read.

## Blocking and deadlocks

A [deadlock](https://yourbasic.org/golang/detect-deadlock/#:~:text=yourbasic.org%2Fgolang,look%20at%20this%20simple%20example.) is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

## Assignment

Run the program. You'll see that it deadlocks and never exits. The `filterOldEmails` function is trying to send on a channel and there is no other goroutine running that can accept the value _from_ the channel.

Fix the deadlock by spawning an anonymous goroutine to write to the `isOldChan` channel instead of using the same goroutine that's reading from it.

### Notes

1. **Goroutines:**

   - In programming, a goroutine is a lightweight thread managed by the Go programming language. It's like a small independent unit of work that can run concurrently with other goroutines.

2. **Queue:**

   - A queue is a data structure that follows the First-In-First-Out (FIFO) principle. Items are added to the back of the queue and removed from the front. Think of it like a line of people waiting for something.

3. **Thread-safe:**

   - This means that the operations performed on the data structure (in this case, the queue) are designed in a way that multiple threads (or goroutines in the context of Go programming) can safely use it without causing unexpected issues or errors.

4. **Typed:**

   - In Go, everything has a type, and channels are no exception. This means that a channel can be specific about the type of data it can transport. For example, you might have a channel of integers or a channel of strings.

5. **Channels:**
   - Channels in Go are a way for goroutines to communicate with each other. They provide a safe and efficient means for one goroutine to send data to another. It's like a pipeline where one goroutine puts something in at one end, and another goroutine takes it out at the other end.

So Channels in Go are like safe communication pipes that allow different parts of a program (goroutines) to easily and safely exchange data. It's a way for them to work together and share information without causing conflicts or issues.

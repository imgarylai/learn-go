package concurrency

// Exercise 6: Concurrency with Goroutines and Channels
//
// This is where Go really shines compared to Node.js!
// Goroutines are like lightweight threads.
// Channels are for communication between goroutines.
// Run tests with: go test -v

import (
	"sync"
	"time"
)

// 1. Basic channel send and receive
// In JS: like resolving a Promise
func ChannelBasics() int {
	// TODO: create an int channel
	// Start a goroutine that sends 42 to the channel
	// Receive from channel and return the value
	return 0
}

// 2. Buffered channel - can hold values without blocking
func BufferedChannel() []int {
	// TODO: create a buffered channel with capacity 3
	// Send 1, 2, 3 to it (no goroutine needed - buffer holds them)
	// Receive all 3 and return as slice
	return nil
}

// 3. Sum numbers using channel
// In JS: similar to Promise.resolve(sum)
func SumWithChannel(nums []int) int {
	// TODO: create channel
	// Start goroutine that calculates sum and sends result
	// Return received sum
	return 0
}

// 4. Channel with range - iterate until closed
// In JS: for await (const item of asyncIterable)
func CollectFromChannel(count int) []int {
	// TODO: create channel
	// Start goroutine that sends 0, 1, 2, ..., count-1 then closes channel
	// Use range to receive all values into slice
	// Hint: close(ch) to signal no more values
	return nil
}

// 5. Select - handle multiple channels (first one wins)
// In JS: Promise.race([promise1, promise2])
func SelectFirst(ch1, ch2 <-chan string) string {
	// TODO: use select to return whichever channel has a value first
	// Hint: select { case v := <-ch1: return v case v := <-ch2: return v }
	return ""
}

// 6. Select with timeout
// In JS: Promise.race([work(), timeout()])
func WithTimeout(work func() int, timeout time.Duration) (int, bool) {
	// TODO: run work() in goroutine, send result to channel
	// Use select with time.After(timeout)
	// Return (result, true) if work completes first
	// Return (0, false) if timeout occurs first
	return 0, false
}

// 7. WaitGroup - wait for multiple goroutines
// In JS: await Promise.all([...])
func SumParallel(slices [][]int) int {
	// TODO: sum each slice in its own goroutine
	// Use sync.WaitGroup to wait for all
	// Use channel to collect partial sums
	// Return total sum
	return 0
}

// 8. Worker pool - limit concurrent workers
// Like limiting concurrent Promise.all to N at a time
func WorkerPool(jobs []int, numWorkers int) []int {
	// TODO: create jobs channel and results channel
	// Start numWorkers goroutines that read from jobs, square the number, send to results
	// Send all jobs to jobs channel, then close it
	// Collect all results
	// Return results (order doesn't matter)
	return nil
}

// 9. Fan-out/Fan-in pattern
// Multiple goroutines read from one channel, results go to one channel
func FanOutFanIn(nums []int, workers int) int {
	// TODO: create input channel with nums
	// Start 'workers' goroutines that each double numbers from input
	// Collect all results and return their sum
	return 0
}

// 10. Mutex - protect shared state
// In JS: you don't usually need this due to single-threaded nature
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	// TODO: safely increment value using mutex
	// Lock, increment, unlock
}

func (c *Counter) Value() int {
	// TODO: safely read value using mutex
	return 0
}

// ConcurrentIncrement tests the Counter
func ConcurrentIncrement(c *Counter, times int) {
	// TODO: start 'times' goroutines, each calling c.Increment()
	// Wait for all to complete
}

// Keep imports used
var _ = sync.WaitGroup{}
var _ = time.Second

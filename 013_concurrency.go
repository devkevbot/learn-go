package main

import (
	"fmt"
	"sync"
	"time"
)

func count(id string) {
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s - %d\n", id, i)
	}
}

func _013GoRoutines() {
	go count("2")
	count("1")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func _013UnbufferedChannels() {
	/* Unbuffered channels block on send and receive until the other
	side is ready. */
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func _013BufferedChannels() {
	/* Send to a buffered channel block only when the buffer is full.
	Receives block when the buffer is empty. */
	ch := make(chan int, 2)
	ch <- 80
	ch <- 19
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonnaci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func _013RangeAndClose() {
	/* Senders can close a channel to indicate that no more values will be sent. */
	c := make(chan int, 10)
	go fibonnaci(cap(c), c)

	/* Range receives values from the channel repeatedly until it is closed. */
	for i := range c {
		fmt.Println(i)
	}
}

func fibonnaci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func _013Select() {
	/* Select lets a goroutine wait on multiple communication
	operations. Select blocks until one of its cases can run. It chooses
	one at random is multiple are ready. */
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonnaci2(c, quit)
}

func _013SelectDefault() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func _013Mutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

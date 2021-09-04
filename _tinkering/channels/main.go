package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numWorkers = 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	mqueue := make(chan string)
	go func() {
		wg.Wait()
		close(mqueue)
	}()

	for id := 1; id <= numWorkers; id++ {
		go func(id int) {
			defer wg.Done()
			sleepFor := time.Duration(id) * time.Second
			time.Sleep(sleepFor)
			mqueue <- fmt.Sprintf("worker %d is done after %d seconds", id, id)
		}(id)
	}

	for msg := range mqueue {
		fmt.Println(msg)
	}
}

package chapter2

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sharedLock sync.Mutex

const runtime = 1 * time.Second

func main() {
	wg.Add(2)
	go greedyWorked()
	go politeWorker()
	wg.Wait()
}

func greedyWorked() {
	defer wg.Done()

	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()
		count++
	}

	fmt.Printf("Greedy worker was able to execute %v work loops.\n", count)
}

func politeWorker() {
	defer wg.Done()

	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		count++
	}

	fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
}

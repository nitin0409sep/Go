package main

import (
	"fmt"
	"sync"
)

// We use mutext - to prevent from race condition while multi threading

type post struct {
	views int
	mu    sync.Mutex // Good practise to define mutex in struct as that will be regarding that struct only
}

// Increment Function
func (p *post) inc(wg *sync.WaitGroup, mutex *sync.Mutex) {

	//  No need to lock the whole function
	// ...... process
	// ...... process
	// ...... process
	// ...... process

	// mutex.Lock()
	p.mu.Lock() // Always put lock on the essentital thing, that matters the most or which should be modified concurrently only
	p.views += 1

	defer func() {
		wg.Done()
		p.mu.Unlock() // Unlock the lock after the work
		// mutex.Unlock()
	}()
}

func main() {
	var wg sync.WaitGroup

	var mutex sync.Mutex

	myPost := post{views: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.inc(&wg, &mutex)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}

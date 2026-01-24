package main

import (
	"fmt"
	"sync"
)

// Go Routines - Light Weight Threads - Usefull at the time when you want to run the things concurrently - Wanna do multi threading

func task(id int, w *sync.WaitGroup) {
	fmt.Println("Task", id)
	defer w.Done() // defer -> runs after function is executed -> just like in react useEffect we have cleaing function
}

// waitGroup in Go Routines - Helpful to know when our all go routines are finished or not

func main() { // Main also runs in a go routines

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg) // go - is used to run go routines
	}

	wg.Wait()

	// time.Sleep(time.Second * 2) // Stop our main function from existing -> So that we can say see our go Routines tasks
}

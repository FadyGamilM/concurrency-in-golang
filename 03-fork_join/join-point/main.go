package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// create a wait group
	waitingGroup := sync.WaitGroup{}

	// main routine requires the number of registered go routines
	waitingGroup.Add(1)

	// run the go routine in anonymous func
	go func() {
		// notify with join point
		defer waitingGroup.Done()
		doSomeWork()
	}()

	// lets see if we execute another go routine we will take the double of the time or less than we are expecting
	waitingGroup.Add(1)
	go func() {
		// notify with join point
		defer waitingGroup.Done()
		doSomeWork()
	}()

	// main routine will wait for all registered waiting groups to be done
	waitingGroup.Wait()

	// calculate the elapsed time
	fmt.Println(time.Since(start))

	fmt.Println("Main routine is done, exit !")
}

func doSomeWork() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done doing some work")
}

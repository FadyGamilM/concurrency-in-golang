package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// define a communication channel btn the main routine nad other go routines
	channel := make(chan struct{})

	// run the go routine in anonymous func
	go func() {
		doSomeWork()
		// fork
		channel <- struct{}{}
	}()

	go func() {
		doSomeWork()
		// fork
		channel <- struct{}{}
	}()

	// notify with join point for 1st fork
	<-channel
	// notify with join point for 2nd fork
	<-channel

	// calculate the elapsed time
	fmt.Println(time.Since(start))

	fmt.Println("Main routine is done, exit !")
}

func doSomeWork() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done doing some work")
}

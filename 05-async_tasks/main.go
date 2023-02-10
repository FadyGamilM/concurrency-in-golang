package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel to communicate
	channel := make(chan struct{})

	// calculate time
	start := time.Now()

	// now pass the channel to each go routine
	go task1(channel)

	go task2(channel)

	go task3(channel)

	go task4(channel)

	// join
	<-channel
	// join
	<-channel
	// join
	<-channel
	// join
	<-channel

	fmt.Println("Time required to execute 4 tasks is => ", time.Since(start))
}

func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1 is done")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task 2 is done")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	fmt.Println("task 3 is done")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("task 4 is done")
	done <- struct{}{}
}

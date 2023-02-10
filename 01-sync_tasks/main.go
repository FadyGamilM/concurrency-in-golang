package main

import (
	"fmt"
	"time"
)

func main() {
	// calculate time
	start := time.Now()

	task1()
	task2()
	task3()
	task4()

	fmt.Println("Time required to execute 4 tasks is => ", time.Since(start))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1 is done")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task 2 is done")
}

func task3() {
	fmt.Println("task 3 is done")
}

func task4() {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("task 4 is done")
}

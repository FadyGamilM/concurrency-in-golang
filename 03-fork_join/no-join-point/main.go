package main

import (
	"fmt"
	"time"
)

func main() {
	// but the execution of the main routine takes only 20 ms
	time.Sleep(20 * time.Millisecond)
	fmt.Println("Main routine is done, exit !")
}

func doSomeWork() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done doing some work")
}

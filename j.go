package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go do(c)

	time.Sleep(4 * time.Second)
	c <- struct{}{}
	fmt.Println("Done")
}

func do(done chan struct{}) {
	for {
		// Do something to stop when done
		fmt.Print(".")
	}
}

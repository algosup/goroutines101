package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go one(c)
	go two(c)

	// Do something with channel here to wait until both one and two have finished

	fmt.Println("We are done")
}

func one(c chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("one")

	// Do something with channel here to signal that you are done
}

func two(c chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("two")

	// Do something with channel here to signal that you are done
}

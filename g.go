package main

import (
	"fmt"
	"time"
)

func main() {
	// Do something to limit the number of goroutines running at the same time to 6
	for {
		go one(c)
	}
}

func one(c chan bool) {
	fmt.Print("+1")
	time.Sleep(5 * time.Second)

	// Do something with channel here to signal that you are done

	fmt.Print("-1")
}

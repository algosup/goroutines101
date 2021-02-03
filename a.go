package main

import (
	"fmt"
	"time"
)

var count int

func main() {
	go add()
	go add()

	time.Sleep(time.Second)
	fmt.Println(count)
}

func add() {
	for i := 0; i < 1e6; i++ {
		count++
	}
}

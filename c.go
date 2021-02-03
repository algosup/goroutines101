package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go add(c)
	go add(c)

	count := <-c + <-c
	fmt.Println(count)
}

func add(c chan int) {
	var local int
	for i := 0; i < 1e6; i++ {
		local++
	}

	c <- local
}

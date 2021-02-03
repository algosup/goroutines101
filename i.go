package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go do(c1)
	go do(c2)
	go do(c3)

	// Pick the fastest channel and display its number of seconds
	// There are no winner if it takes more than 6 seconds
	fmt.Println("The winner is XX with", v, "secondes")
	fmt.Println("No winner, it took more than 6 seconds")
}

func do(c chan int) {
	t := 2 + rand.Intn(15)
	time.Sleep(time.Duration(t) * time.Second)
	c <- t
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := make([]int, 1e6)
	fmt.Println("Generating numbers...")
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(1e5)
	}

	c := make(chan int)
	fmt.Println("Starting calculation...")
	start := time.Now()
	go sum(numbers, c)
	fmt.Println("Total value:", <-c)
	fmt.Println("Sum took:", time.Since(start))
}

func sum(input []int, output chan int) {
	// Modify this code to make it able to run in parallel (to speed things up)
	r := 0
	for _, v := range input {
		if isPrime(v) {
			r += v
		}
	}
	output <- r
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

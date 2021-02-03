package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var lock sync.RWMutex

func main() {
	go add()
	go add()

	time.Sleep(time.Second)
	lock.Lock()
	fmt.Println(count)
	lock.Unlock()
}

func add() {
	for i := 0; i < 1e6; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
}

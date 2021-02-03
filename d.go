package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go write(c)
	read(c)
}

func write(c chan struct{}) {
	fmt.Println("1 - Started writing to channel", time.Now().Format("15:04:05"))
	c <- struct{}{}
	fmt.Println("2 - Finished writing to channel", time.Now().Format("15:04:05"))
	fmt.Println("3 - Waiting for 3 seconds", time.Now().Format("15:04:05"))
	time.Sleep(3 * time.Second)
	fmt.Println("4 - Started writing to channel", time.Now().Format("15:04:05"))
	c <- struct{}{}
	fmt.Println("5 - Finished writing to channel", time.Now().Format("15:04:05"))
}

func read(c chan struct{}) {
	fmt.Println("A - Waiting for 2 seconds", time.Now().Format("15:04:05"))
	time.Sleep(2 * time.Second)
	fmt.Println("B - Started reading from channel", time.Now().Format("15:04:05"))
	<-c
	fmt.Println("C - Finished reading from channel", time.Now().Format("15:04:05"))
	fmt.Println("D - Started reading from channel", time.Now().Format("15:04:05"))
	<-c
	fmt.Println("E - Finished reading from channel", time.Now().Format("15:04:05"))
}

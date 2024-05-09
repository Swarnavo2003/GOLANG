package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, world!")
	// time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHello function ended successfully")
}

func sayHi() {
	fmt.Println("Hi Swarnavo :)")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hi Swarnvo Function ended;)")
}

func main() {
	fmt.Println("Learning Goroutines...")

	go sayHello()
	go sayHi()

	time.Sleep(1000 * time.Millisecond)
}
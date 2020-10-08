package main

import "fmt"

func someFunc(channel chan int) {
	channel <- 5
}

func main() {
	channel := make(chan int, 6)

	go someFunc(channel)

	number := <-channel
	fmt.Print(number)
}
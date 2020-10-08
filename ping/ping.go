package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	for {
		channel <- "ping"
		fmt.Println("Foo sends ping")

		pong := <-channel
		fmt.Println("Foo receives: ", pong)
	}
}

func bar(channel chan string) {
	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
	for {
		ping := <-channel
		fmt.Println("bar receives:", ping)

		channel <- "pong"
		fmt.Println("bar sends pong")
	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	messages := make(chan string)

	go foo(messages) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(messages)


	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}

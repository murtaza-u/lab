package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // unbuffered bi-di channel
	go send(ch)
	recv(ch)
}

// bi-di channel is implicitly typecasted send-only channel
func send(ch chan<- int) {
	for {
		ch <- 0
		time.Sleep(time.Second)
	}
}

// bi-di channel is implicitly typecasted receive-only channel
func recv(ch <-chan int) {
	for {
		n := <-ch
		fmt.Println(n)
	}
}

package main

import (
	"fmt"
)

func send(send chan<- int, arr []int) {
	for _, v := range arr {
		send <- v
	}
	close(send)
}

func double(out <-chan int, in chan<- int) {
	for x := range out {
		x *= 2
		in <- x
	}
	close(in)
}

func printFromCh(out <-chan int, done chan bool) {
	for {
		x, ok := <-out
		if !ok {
			break
		}
		fmt.Println(x)
	}
	done <- true
}

//-------------------- main ------------------------------
func main() {

	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch1, ch2 := make(chan int), make(chan int)
	done := make(chan bool)

	go send(ch1, arr)
	go double(ch1, ch2)
	go printFromCh(ch2, done)

	fmt.Println("done:", <-done)
}

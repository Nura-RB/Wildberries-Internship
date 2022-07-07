package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func printSmth(in chan<- int) {
	count := 1
	for {
		in <- count
		count++
	}
}

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	go printSmth(ch)
	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		case <-ctx.Done():
			fmt.Println("timeout")
			os.Exit(1)
		}
	}
}

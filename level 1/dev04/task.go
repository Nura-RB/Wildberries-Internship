package main

import "fmt"

func worker(ch <-chan interface{}) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	var n int
	ch := make(chan interface{})

	fmt.Println("Enter numbers of workers")
	fmt.Scanln(&n)

	data := 0
	for i := 0; i < n; i++ {
		go worker(ch)
		for {
			ch <- data
			data++
		}
	}

	close(ch)
}

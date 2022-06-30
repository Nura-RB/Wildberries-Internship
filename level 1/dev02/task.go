package main

import "fmt"

func printSqr(ch <-chan int) {
	for {
		x := <-ch
		if x == 0 {
			break
		}
		fmt.Println(x * x)
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	go printSqr(ch)

	for i := 0; i < len(arr); i++ {
		ch <- arr[i]
	}

	ch <- 0
}

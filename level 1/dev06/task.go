package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 1) with additional channel
func writer(ch chan<- int, end <-chan bool) {
	for i := 0; ; i++ {
		select {
		case <-end:
			return
		default:
			ch <- i
		}
	}
}

func waiter(ch chan<- bool, n int) {
	// fmt.Println("strat")
	time.Sleep(time.Duration(n) * time.Second)
	// fmt.Println("end")
	ch <- true
}

// 2) sending signal by pointer

func f(flag *bool) {
	for {
		if *flag {
			fmt.Println("working")
		} else {
			break
		}
	}
	fmt.Println("stop")
}

// 3) close by WaitGrop
func tryCloseFirst() {
	ch := make(chan struct{})
	wg := &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		for {
			select {
			case <-ch:
				return
			}
		}
	}()
}

// 4) close by close()
func tryCloseSecond() {
	ch := make(chan struct{})
	wg := &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		close(ch)
	}()
}

// 5) with context

func tryCloseThird() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	child := func(ctx context.Context) <-chan int {
		inc := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case ch <- inc:
					inc++
					fmt.Println("POVYSHEN")
				}
			}
		}()
		return ch
	}

	for i := range child(ctx) {
		if i == 5 {
			break
		}
	}
}

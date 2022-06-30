package main

import (
	"fmt"
	"sync"
)

func calcSum(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, v := range arr {
		sum += v * v
	}
	fmt.Println(sum)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(1)
	go calcSum(arr, &wg)

	wg.Wait()

}

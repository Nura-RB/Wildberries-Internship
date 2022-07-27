package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func binarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == item {
			return mid
		}
		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := make([]int, 100)

	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(200)
	}

	sort.Ints(arr)

	fmt.Println(arr)
	fmt.Println(binarySearch(arr, 45))
}

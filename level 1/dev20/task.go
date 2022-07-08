package main

import (
	"fmt"
	"os"
)

func main() {
	arr := make([]string, len(os.Args)-1)
	j := 0
	for i := len(os.Args) - 1; i > 0; i-- {
		arr[j] = os.Args[i]
		j++
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		if i != len(arr)-1 {
			fmt.Print(" ")
		}
	}
}

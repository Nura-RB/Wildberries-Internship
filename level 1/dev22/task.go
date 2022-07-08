package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int
	var operator string

	fmt.Scanln(&a)
	fmt.Scanln(&operator)
	fmt.Scanln(&b)

	switch operator {
	case "+":
		fmt.Println(a.Add(&a, &b))
	case "-":
		fmt.Println(a.Sub(&a, &b))
	case "*":
		fmt.Println(a.Mul(&a, &b))
	case "/":
		c := big.NewInt(0)
		if b.Cmp(c) == 0 {
			fmt.Println("division by 0!")
			break
		}
		fmt.Println(a.Div(&a, &b))
	}
}

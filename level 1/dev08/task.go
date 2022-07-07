package main

import (
	"fmt"
	"strconv"
)

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func main() {
	var n int64
	var i int64
	var val int64

	fmt.Println("Введите число:")
	fmt.Scanf("%d", &n)

	fmt.Println("Введите бит, который нужно преобразовать:")
	fmt.Scanf("%d", &i)

	fmt.Println("Введите 1 или 0:")
	fmt.Scanf("%d", &val)
	nString := strconv.FormatInt(n, 10)
	v, _ := ConvertInt(nString, 10, 2)
	fmt.Println("До преобразования:", v)

	// Работает только с преобразованием 0 в 1
	if val == 1 {
		n |= val << i
	} else {
		n &= n ^ val<<i - 1
	}

	nString = strconv.FormatInt(n, 10)
	v, _ = ConvertInt(nString, 10, 2)

	fmt.Println("После преобразования:", v)
	fmt.Println("Результат:", n)
}

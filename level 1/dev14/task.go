package main

import "fmt"

// type switch
func someFunc(data interface{}) {
	switch t := data.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case chan interface{}:
		fmt.Println("channel")
	default:
		fmt.Printf("%T\n", t)
	}
}

func main() {
	ch := make(chan interface{})
	someFunc(ch)
}

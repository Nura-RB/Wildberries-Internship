package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) sayHello() {
	fmt.Println("Hello")
}

type Action struct {
	Human
}

func main() {
	act := Action{}
	act.sayHello()

}

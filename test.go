package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	n, error := fmt.Println("this is a test")
	fmt.Println(n, error)
	foo()
	a = 1
	b = 2

	a = int(b)
	fmt.Println(a)
	fmt.Println(b)
}

func foo() {
	fmt.Println("this is sam")
}

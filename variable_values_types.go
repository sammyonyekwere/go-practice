package main

import (
	"fmt"
)

// Package level scoped variable
var a, y int
var u int = 42
var v string = "James Bond"
var w bool = true

// Go underlying type
type saamz int

var games, b saamz

func main() {
	n, error := fmt.Println("this is a test")
	fmt.Println(n, error)
	foo()
	a = 1
	b = 2

	games = 42
	fmt.Println(games)
	fmt.Printf("%T", games)

	// Type conversion
	y = int(games)
	fmt.Println(y)
	fmt.Printf("%T", y)

	a = int(b)
	fmt.Println(a)
	fmt.Println(b)

	x := 42
	y := "James Bond"
	z := true

	// Multiple print statement
	fmt.Println(x)
	fmt.Println(y)

	// Single Print Statement
	fmt.Println(x, y, z)

	fmt.Println(u, v, w)

	s := fmt.Sprintf("%d\t%s\t%v", u, v, w)

	fmt.Println(s)

}

func foo() {
	fmt.Println("this is sam")
}

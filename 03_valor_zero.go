package main

import "fmt"

var a int
var b float32
var c string

func main() {

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)

}

/*
	InputTerminal:
 	a: 0, int
	b: 0, float32
	c: , string
*/

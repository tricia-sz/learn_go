package main

import "fmt"

type hotdog int

var b hotdog
var c int

func main() {

	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}

/*
	InputTerminal:
	main.hotdog
	int
*/

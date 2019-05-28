package main

import "fmt"

func main() {
	f := func(x int) int {
		return x
	}
	x := foo(f, 1234)
	fmt.Println(x)
}

func foo(f func(x int) int, i int) int {
	g := f(i)
	return g
}
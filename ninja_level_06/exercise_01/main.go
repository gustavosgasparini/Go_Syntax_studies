package main

import "fmt"

func main()  {
	fmt.Println(foo())
	fmt.Println(bar())

	x := foo()
	y, s := bar()

	fmt.Println(x, y, s)
}

func foo() int {
	return 1234
}

func bar() (int, string) {
	return 1234, "wtf duuuude"
}
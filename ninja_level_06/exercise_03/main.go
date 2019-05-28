package main

import "fmt"

func main()  {
	defer fmt.Println(foo())
	fmt.Println(bar())
}

func foo() string {
	return "Last but not least - defer"
}

func bar() string {
	return "first but last - no defer"
}

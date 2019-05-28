package main

import "fmt"

func main() {
	ii1 := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(foo(ii1...))

	ii2 := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(bar(ii2))
}

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar(a []int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
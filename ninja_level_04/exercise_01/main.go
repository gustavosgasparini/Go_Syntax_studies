package main

import "fmt"

func main()  {
	x := [5]int{13,24,35,67,43}
	for i, v := range x {
		fmt.Println("index",i,"=",v)
	}
	fmt.Printf("The array is of type: %T\n", x)
}

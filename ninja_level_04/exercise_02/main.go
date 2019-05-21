package main

import "fmt"

func main()  {
	x := []int{01,04,05,10,12,16,20,22,29,33}
	for i, v := range x {
		fmt.Println("index",i,"=",v)
	}
	fmt.Printf("Slice of type: %T\n", x)
}

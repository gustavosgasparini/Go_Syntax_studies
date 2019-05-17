package main

import "fmt"

func main()  {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v divided by 4, modulus is %v\n", i, i%4)
	}
}

package main

import "fmt"

func main()  {
	f := func() {
		fmt.Println("assigned to a variable")
	}

	f()
	fmt.Printf("%T\n", f)
}

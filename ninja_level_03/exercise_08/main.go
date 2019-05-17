package main

import "fmt"

func main()  {

	switch {
	case false:
		fmt.Println("This case is false.") // It won't be printed because it's false
	case true:
		fmt.Println("This case is true") // It will be printed
	}

}

package main

import "fmt"

func main()  {

	birth := 1990
	for {
		if birth > 2019 {
			break
		}
		fmt.Println(birth)
		birth++
	}

}
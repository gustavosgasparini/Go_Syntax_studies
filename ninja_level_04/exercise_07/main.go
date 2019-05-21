package main

import "fmt"

func main()  {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)

	mm := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(mm)

	x := [][]string{jb,mm}
	fmt.Println(x)

	for rec, char := range x {
		fmt.Println("record:",rec)
		for i, v := range char {
			fmt.Println("index",i,"=",v)
		}
	}
}

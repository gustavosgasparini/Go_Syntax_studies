package main

import "fmt"

type person struct {
	firstName string
	lastName string
	iceCreamFav []string
}

func main()  {
	p1 := person{
		firstName: "Gustavo",
		lastName: "Gasparini",
		iceCreamFav: []string{
			"chocolate",
			"strawberry",
			"vanilla",
		},
	}
	fmt.Println(p1)
	for i, v := range p1.iceCreamFav {
		fmt.Println(i,v)
	}

	fmt.Println("-------------------")

	p2 := person{
		firstName: "Rebeca",
		lastName: "Rancan",
		iceCreamFav: []string{
			"Choolate",
			"Strawberry",
			"Passion fruit",
		},
	}
	fmt.Println(p2)
	for i, v := range p2.iceCreamFav {
		fmt.Println(i,v)
	}

	fmt.Println("-------------------")

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Printf("%s %s\n", v.firstName, v.lastName)
		for _, val := range v.iceCreamFav {
			fmt.Println("Ice cream:",val)
		}
		fmt.Println("--------------------")
	}
}
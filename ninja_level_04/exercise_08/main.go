package main

import "fmt"

func main()  {
	m := map[string][]string {
		"bond_james":	   []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": 		   []string{"Being evil", "Ice cream", "sunsets"},
	}
	fmt.Println(m)

	for rec, char := range m {
		fmt.Println("---------------")
		fmt.Println("record:",rec)
		for i,v := range char {
			fmt.Println("Needs",i,"=",v)
		}
	}
}

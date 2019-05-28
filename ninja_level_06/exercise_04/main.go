package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old")
}

func main() {
	p1 := Person{
		first: "Gustavo",
		last:  "Gasparini",
		age:   29,
	}

	p1.speak()
}

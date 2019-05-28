package main

import "fmt"

func main() {
	p := person{
		first: "Gustavo",
		last:  "Gasparini",
		age: 29,
	}

	a := assassin{
		first:"Rebeca",
		last:"Rancan",
		bodyCount:14,
	}

	p.normalPerson()
	a.killer()
}

type person struct {
	first string
	last  string
	age   int
}

type assassin struct {
	first     string
	last      string
	bodyCount int
}

func (a assassin) killer() {
	fmt.Println("my name is", a.first, a.last, "and i killed", a.bodyCount, "humanbeigns")
}

func (p person) normalPerson() {
	fmt.Println("my name is", p.first, p.last, "and i'm", p.age, "years old, just a normal person")
}

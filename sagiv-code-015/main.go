package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("i am", p.first)
}

func main() {

	p1 := person{
		first: "sagiv",
	}

	p2 := person{
		first: "gal",
	}

	p1.speak()
	p2.speak()

}

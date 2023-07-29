package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secreteAgent struct {
	person
	LTK   bool
	first string
}

func main() {
	sa1 := secreteAgent{
		person: person{
			first: "sagiv",
			last:  "oron",
			age:   27,
		},
		LTK:   true,
		first: "Galgalon",
	}

	fmt.Printf("%#v\n", sa1)
	fmt.Println("----------------------------------------")

	p2 := person{
		first: "gal",
		last:  "ohayon",
		age:   25,
	}

	fmt.Printf("the type of %v is %T or %#v\n", sa1, sa1, sa1)
	fmt.Printf("the type of %v is %T or %#v\n", p2, p2, p2)

	fmt.Println("----------------------------------------")

	fmt.Printf("we can also display it like this - %v %v %v,%v\n", sa1.first, sa1.last, sa1.age, sa1.LTK)
	fmt.Println("----------------------------------------")
	fmt.Printf("we can also display it like this - %#v %#v %#v,%#v\n", sa1.first, sa1.last, sa1.age, sa1.LTK)
	fmt.Println("----------------------------------------")
	fmt.Printf("we can also display it like this - %T %T %T,%T\n", sa1.first, sa1.last, sa1.age, sa1.LTK)
	fmt.Println("----------------------------------------")
	fmt.Println(sa1.person, sa1.LTK)
	fmt.Println("----------------------------------------")
	fmt.Println(sa1.first, sa1.first)
	fmt.Println("----------------------------------------")
	fmt.Println(sa1.first, sa1.person.first)
}

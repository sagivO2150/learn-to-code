package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Println(a)

	b := []string{"sagiv", "gal", "mia"}
	fmt.Println(b)

	var c [4]int
	c[0] = 4
	c[1] = 5
	c[2] = 6
	c[3] = 7
	fmt.Println(c)

	fmt.Printf("the type of a is - %T\n", a)
	fmt.Printf("the type of a is - %T\n", c)

	c = a
	fmt.Println(c)
}

package main

import (
	"fmt"

	"github.com/sagivO2150/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(puppy.From13())
	//puppy.From13()

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)

	//option num 2

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}

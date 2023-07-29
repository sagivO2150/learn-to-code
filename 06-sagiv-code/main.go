package main

import "fmt"

func main() {
	// xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(xi)
	// fmt.Println("---------------")

	// xi = append(xi[:4], xi[5:]...)
	// fmt.Println(xi)
	// fmt.Println("---------------")

	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("------------------------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println(xi)
	fmt.Println("------------------------------")

	// for i := 0; i < cap(xi); i++ {
	// 	fmt.Printf("the vlaue of xi in place %v is - %v\n", i, xi[i])
	// }

}

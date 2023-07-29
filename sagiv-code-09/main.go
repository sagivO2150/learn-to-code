package main

import (
	"fmt"
)

func main() {
	// am := map[string]int{
	// 	"Todd":  42,
	// 	"Henry": 16,
	// 	"sagiv": 27,
	// }

	// am["mom"] = 47

	// // fmt.Println("the age of sagiv is - ", am["sagiv"])
	// fmt.Println("before delete", am)
	// // fmt.Printf("%#v\n", am)
	// // fmt.Printf("%T\n", am)
	// // fmt.Println(len(am))

	// delete(am, "Todd")
	// // fmt.Println("--------------------------")
	// //for range over map
	// fmt.Println("after delete -")
	// for k, v := range am {
	// 	fmt.Println(k, v)
	// }

	// for _, v := range am {
	// 	fmt.Println(v)
	// }

	// for k := range am {
	// 	fmt.Println(k)
	// }
	// 	fmt.Println("--------------------------")

	// 	//for range over slice
	// 	xi := []int{42, 43, 44}

	// 	for k, v := range xi {
	// 		fmt.Println(k, v)
	// 	}

	// 	for _, v := range xi {
	// 		fmt.Println(v)
	// 	}

	// 	for k := range xi {
	// 		fmt.Println(k)
	// 	}

	an := make(map[string]int)
	an["gal"] = 25
	an["dad"] = 60
	an["mom"] = 50

	v, ok := an["mom"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("mon didnt exist")

	}

	if z, ok := an["dan"]; !ok {
		fmt.Println("dan doesnt exist")
	} else {
		fmt.Println("this is the age of dad -", z)
	}

	if r, ok := an["gal"]; ok {
		fmt.Println("this is the age of gal -", r)
	} else {
		fmt.Println("gag doesnt exist")
	}

	fmt.Println("this is the age of mom-", v)
	// fmt.Println(an)
	// fmt.Println(len(an))
	// delete(an, "mom")
	// delete(an, "mon")

	// fmt.Println("wonder what will happen?")
	// fmt.Println(an["mon"])

	// fmt.Println(len(an))
	// fmt.Printf("%#v\n", an)
	// fmt.Printf("%T\n", an)
}

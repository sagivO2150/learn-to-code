package main

import "fmt"

func main() {
	foo()
	bar("Todd")
	s := aloha("kuka")
	fmt.Println(s)

	s1, n := dogYears("Todd", 5)
	fmt.Println(s1, n)

}

func foo() {
	fmt.Print("i am from foo\n")
}

func bar(s string) {
	fmt.Println("my name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("this call me mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old many dog years "), age
}

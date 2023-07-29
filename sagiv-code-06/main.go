package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "martini", "chocolate"}
	jm := []string{"Jenny", "Moneypanny", "ginuse", "wolverin traks"}
	jr := []string{"kaki", "pipi", "bulbul", "tahat"}
	fmt.Println(jb)
	fmt.Println(jm)

	xss := [][]string{jb, jm}
	xss = append(xss, jr)
	fmt.Println(xss)

}

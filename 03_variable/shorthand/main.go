package main

import "fmt"

var zewnatrz string = "deklaracja poza funkcja"

func main() {
	//wersja skrócona wewnątrz funkcji
	a := "string"
	b := 2
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println("%v \n", zewnatrz)
}

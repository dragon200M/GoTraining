package main

import "fmt"

func main() {

	var a string
	var b int
	var c float32
	var d bool

	fmt.Println("Wartość domyślna:")
	fmt.Printf("String: %q \n", a)
	fmt.Printf("integer: %v \n", b)
	fmt.Printf("float: %v \n", c)
	fmt.Printf("bool: %v \n\n", d)

	fmt.Println("Typ:\n")
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}

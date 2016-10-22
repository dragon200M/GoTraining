package main

import "fmt"

func main() {
	var name string

	fmt.Scan(&name)

	switch name {

	case "Piotr":
		fmt.Println("hello my friend ", name)
	case "Łucja":
		fmt.Println("hello my wife")
	default:
		fmt.Println("end")
	}

}

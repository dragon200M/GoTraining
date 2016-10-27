package main

import "fmt"

func main() {

	if true || false {
		fmt.Println("This ran")
	}

	if true && false {
		fmt.Println("Nie działa")
	}

	if !false {
		fmt.Println("Działa")
	}
}

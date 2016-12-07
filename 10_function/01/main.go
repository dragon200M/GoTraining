package main

import "fmt"

func main() {
	greeting("Maciej")
	greeting("Łucja")

	twoparam("Maciej", "Michalik")

}

func greeting(name string) {
	fmt.Println("Hello ", name)

}

func twoparam(fname, lname string) {
	fmt.Println(fname, lname)
}

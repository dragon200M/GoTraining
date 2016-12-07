package main

import "fmt"

const metersToFeets float64 = 3.2808399

func main() {

	a := 43
	fmt.Println("a - ", a)
	fmt.Printf("a's memory address %d \n", &a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b) //43

	var meters float64
	fmt.Print("Enter meters swam:")
	fmt.Scan(&meters)
	feet := meters * metersToFeets

	fmt.Println(meters, " meters is", feet, " feets")
}

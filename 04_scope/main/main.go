package main

import "fmt"
import "github.com/dragon200M/GoTraining/04_scope/vis"

var x int = 42

func main() {
	fmt.Println(vis.LastName)
	vis.PrintVa()
	fmt.Println()
	{
		fmt.Println(x)
		y := 34
		fmt.Println(y)

	}
	//fmt.Println(y) y niedostepne

	fmt.Println()
	fmt.Println(increment())
	fmt.Println(increment())

	//anonymous function

	inc := func() int {
		x++
		return x
	}

	fmt.Printf("\n anonymous function \n")
	fmt.Println(inc())
	fmt.Println(inc())

	inc2 := wrapper()
	fmt.Printf("\n Return function \n")
	fmt.Println(inc2())
	fmt.Println(inc2())

}

func increment() int {
	x++
	return x
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}

}
/*
Types:
Method sets
Boolean types
Numeric types
String types
Array types
Slice types
Struct types
Pointer types
Function types
Interface types
Map types
Channel types
*/
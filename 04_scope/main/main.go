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
}


func increment() int {
	x++
	return x
}

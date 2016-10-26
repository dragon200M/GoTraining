package main

import "fmt"

func main() {
	//1
	t := average(2, 4, 5, 7)
	fmt.Println(t)

	a := []float64{2, 4, 5, 7}
	//2
	fmt.Println(average(a...))

}

func average(n ...float64) float64 {

	var total float64

	for _, i := range n {

		total += i
	}

	return total / float64(len(n))

}

package main

import "fmt"

func main() {
	//1
	t := average(2, 4, 5, 7)
	fmt.Println("Total:", t)

	a := []float64{2, 4, 5, 7}
	//2
	fmt.Println("Total:", average(a...))

	//3
	b := []float64{2, 4, 5, 7}
	fmt.Println("Total:", average2(b))
	//4
	fmt.Println("Total:", average2([]float64{2, 4, 5, 7}))

}

func average(n ...float64) float64 {

	var total float64

	for k, i := range n {
		fmt.Print(k, ":", i, " ")
		total += i
	}

	return total / float64(len(n))

}

func average2(n []float64) float64 {

	var total float64

	for k, i := range n {
		fmt.Print(k, ":", i, " ")
		total += i
	}

	return total / float64(len(n))

}

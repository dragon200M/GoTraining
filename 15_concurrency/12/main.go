package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for a := range factorial2(11) {
		fmt.Println(a)
	}
	w := factorial(11)

	fmt.Println(w)
	w1 := factorial(11)

	fmt.Println(w1)
	stop := time.Since(start)

	fmt.Println(stop)

}

func factorial(n int) int {
	total := 1

	for i := n; i > 0; i-- {

		total *= i
	}

	return total

}

func factorial2(n int) chan int {
	c := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {

			total *= i
		}
		c <- total
		close(c)
	}()

	return c

}
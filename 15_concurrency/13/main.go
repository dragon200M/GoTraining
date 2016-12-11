package main

import "fmt"

func main() {
	//pipeline
	zm := gen(2, 3, 4)

	for n := range seq(zm) {
		fmt.Println(n)
	}

}

func gen(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {

			out <- n
		}
		close(out)

	}()

	return out
}

func seq(c chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range c {

			out <- n * n
		}
		close(out)

	}()

	return out
}

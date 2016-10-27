package main

import "fmt"

func vis(numbers []int, call func(int)) {

	for _, n := range numbers {
		call(n)
	}
}

func factorial(x int) int {

	if x == 0 {
		return 1
	}

	return x * factorial(x-1)

}

func main() {

	vis([]int{1, 2, 3, 4}, func(n int) {

		fmt.Println(n)
	})

	fmt.Println(factorial(4))

}

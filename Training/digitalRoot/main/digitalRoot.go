package main

import "strconv"
import "fmt"

func digitalRoot(n int) int {
	a := strconv.Itoa(n)
	sum := 0

	for _, value := range a {

		i, err := strconv.ParseInt(string(value), 0, 64)

		if err != nil {
			panic(err)
		}
		sum += int(i)
	}

	if sum < 10 {
		return sum
	} else {
		return digitalRoot(sum)

	}

}

func main() {

	fmt.Println(digitalRoot(942))

}

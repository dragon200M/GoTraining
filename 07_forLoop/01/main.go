package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		i++
		if i%2 == 0 {
			fmt.Println(i)
			continue

		}

		if i >= 20 {
			break
		}

	}
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))

	}

}

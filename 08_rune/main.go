package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println([]byte("Hello"))

	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))

	}

	for i := 5000; i <= 5100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))

	}

}

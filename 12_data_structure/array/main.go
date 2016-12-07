package main

import "fmt"

func main() {

	var ar [58]string

	for i := 65; i <= 122; i++ {

		ar[i-65] = string(i)
	}

	fmt.Println(ar)

	var x [256]byte

	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)

		if i >= 50 {
			break
		}
	}

	for _, v := range ar {

		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}

}

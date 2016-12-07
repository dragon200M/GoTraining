package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a) //0xc82000a3b8

	var b *int = &a
	fmt.Println(b) //0xc82000a3b8

	fmt.Println(*b) //43

	*b = 1
	fmt.Println(a)  //1
	fmt.Println(&a) //0xc82000a3b8

}

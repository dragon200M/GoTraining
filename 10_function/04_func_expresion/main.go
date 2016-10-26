package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello")
	}
	greeting()

	greet := makegreeter()
	fmt.Println(greet) //0x4011e0
	fmt.Println(greet()) //hello world
	fmt.Printf("%T\n",greet)//func() string
	fmt.Printf("%T\n",greet())//string
}

func makegreeter() func() string {

	return func() string {
		return "hello world"
	}
}

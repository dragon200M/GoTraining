package main

import "fmt"

const p string = "stała p"

const (
	F = iota //0
	G = iota //1
	H = iota //2

	L = "Go" //iota 4
	I = iota //4

)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func main() {

	const q = "stała q"

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(L)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)

	{

		fmt.Println(q)
		fmt.Println(p)
	}
	cos()

	fmt.Printf("binary\t\tdecimal\n")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)

	x := 42
	t := x + 19
	fmt.Println(t)

}

func cos() {

	fmt.Println(p)

}

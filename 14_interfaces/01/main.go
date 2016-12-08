package main

import "fmt"

type Square struct {
	a float64
}

func (z Square) area() float64 {

	return z.a * z.a
}

func (z Square) circuit() float64 {
	return 4*z.a
}

type Triangle struct {
	a float64
	h float64
}



func (t Triangle) area() float64 {

	return (t.a * t.h) / 2
}

func (t Triangle) circuit() float64 {
	return 0
}

func info(s Shape) {
	fmt.Printf("%T \n",s)
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.circuit())

}

type Shape interface {
	area() float64
	circuit() float64
}

func main() {
	s := Square{12.4}
	t := Triangle{3.42, 8.2}


	info(s)
	info(t)

	const con = 42
	fmt.Println(con+2.1) //work

	const  con2 int =  42
	fmt.Println(con2+3.1)//doesn't work


}

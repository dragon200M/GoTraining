package main

import "fmt"

type Square struct {
	a float64
}

func (z Square) area() float64 {

	return z.a * z.a
}

type Traingle struct {
	a float64
	h float64
}

func (t Traingle) area() float64 {

	return (t.a * t.h) / 2
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())

}

type Shape interface {
	area() float64
}

func main() {
	s := Square{12.4}
	t := Traingle{3.42, 8.2}

	info(s)
	info(t)

}

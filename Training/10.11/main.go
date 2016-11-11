package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(odw(12560))

	tab := []string{"abc", "def", "ghi"}

	fmt.Println(odwSl(tab))

	fib(5)
	fmt.Println()
	fmt.Println(fibR(5))

	fmt.Println(ul(8811))
	fmt.Println(ul2(8811))
}

//odwrócenie integera

func odw(liczba int) string {

	var tmp int
	nl := ""

	for liczba > 0 {
		tmp = liczba % 10
		liczba = liczba / 10

		nl = nl + strconv.Itoa(tmp)
	}

	return nl

}

//odwracanie słów w tablicy

func odwSl(sl []string) []string {

	d := len(sl) - 1
	for i := 0; i < d/2; i++ {
		tmp := sl[i]
		sl[i] = sl[d-i]
		sl[d-i] = tmp
	}

	return sl
}

//fibonacci-ego - rekurencyjnie lub iteracyjnie
func fib(n int) {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		fmt.Print(b, " ")
		b += a
		a = b - a

	}

}

func fibR(n int) int {

	if n < 3 {
		return 1
	}

	return fibR(n-1) + fibR(n-2)
}

//ułamek z cyfr liczby większy od 7

func ul(n int) (float32, int) {
	t1 := n
	t := n
	li := 0
	for t > 0 {
		t = t / 10
		li++

	}

	var u float32
	u = 0
	for n > 0 {
		u = u + float32(n%10)

		n = n / 10

	}

	w := u / float32(li)

	if w > 7 {
		return w, li
	}

	if li > 0 {
		return ul(t1 / 10)
	}

	return 0.0, 0
}

func ul2(n int) (float32, int) {
	t1 :=n
	t := n
	li := 0
	for t > 0 {
		t = t / 10
		li++

	}

	var u float32
	u = 0

	tl := li
	for j := 0; j < tl; j++ {



		for n > 0{

			u = u + float32(n%10)
			n = n / 10

		}

		if u / float32(li) > 7.0 {

			return u / float32(li),li

		}

		n = t1/10
		t1 = t1/10
		u = 0
		li--


	}



	return 0,0
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(odw(12560))

	tab := []string{"abc","def","ghi"}

	fmt.Println(odwSl(tab))

	fib(5)
	fmt.Println()
	fmt.Println(fibR(5))
}


//odwrócenie integera

func odw(liczba int) string {

	var tmp int
	nl :=""

	for liczba > 0 {
		tmp = liczba % 10
		liczba = liczba / 10

		nl = nl + strconv.Itoa(tmp)
	}

	return nl

}


//odwracanie słów w tablicy

func odwSl(sl []string) []string{

	d:=len(sl)-1
	for i:=0;i< d/2;i++{
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

	for i:=0;i<n;i++{
		fmt.Print(b," ")
		b +=a
		a = b-a

	}

}

func fibR(n int) int {


	if n < 3 {
		return 1
	}

	return fibR(n-1)+fibR(n-2)
}

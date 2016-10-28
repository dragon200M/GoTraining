package main


import (
	"fmt"
	"strconv"
)

func odw(c string) string {
	r := []rune(c)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]

	}

	return string(r)
}

func odw2(c string) string {
	r := []rune(c)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[i]
		r[i] = r[len(r)-1-i]
		r[len(r)-1-i] = tmp

	}

	return string(r)
}

func bez(n int) int {

	if n > 0 {
		return n
	} else if n < 0 {

		return n * (-1)
	}
	return 0
}

func sum(s []float64) float64 {

	var wynik float64 = 0

	for i := 0; i < len(s); i++ {
		wynik += s[i]

	}

	return wynik

}

func sum2(s []float64) float64 {

	var wynik float64 = 0.0
	for _, i := range s {
		wynik += i

	}
	return wynik

}

func count(tab []int, s int) int {

	wynik := 0
	for _, i := range tab {
		if i == s {
			wynik += 1
		}
	}

	return wynik
}

func count2(tab []int) []int {

	var tmp = make([]int, len(tab))

	prev := tab[0]
	c := 0
	for i := 0; i < len(tab); i++ {

		if tab[i] == prev{
			tmp[c] +=  1

		}else{
			c++
			prev = tab[i]
			tmp[c] = 1

		}
	}
	w := 0
	for k := 0; k < len(tmp); k++ {
		if tmp[k] != 0 {
			w +=1
		}
	}

	var t =make([]int,w)

	for l := 0; l < w; l++ {

		t[l] = tmp[l]

	}


	return t
}

func minMax(tab []int) string{

	max := tab[0]
	min := tab[0]
	for i:=1; i<len(tab);i++ {

		if tab[i] > max {
			max =tab[i]
		}
		if tab[i] < min {
			min =tab[i]
		}


	}

	str :="Max:"+strconv.Itoa(max)+" Min:"+ strconv.Itoa(min)

	return str
}

func main() {
	fmt.Println("Odwrotność: ")
	fmt.Println(odw("abcd"), "-> dcba")
	fmt.Println(odw("abcde"), "-> edcba")
	fmt.Println(odw2("abcd"), "-> dcba")
	fmt.Println(odw2("abcde"), "-> edcba")

	fmt.Println("Wartość bezwzględna: ")
	fmt.Println(bez(-1))
	fmt.Println(bez(-(-1)))
	fmt.Println(bez(8))
	fmt.Println(bez(-8))

	fmt.Println("Suma: ")
	fmt.Println(sum([]float64{1, 2, 3, -1, 45}))
	fmt.Println(sum2([]float64{1, 2, 3, -1, 45}))

	fmt.Println("Ilość elementów: ")
	fmt.Println(count([]int{4, 1, 6, 1, 1, 1, 1}, 1))

	fmt.Println("Różne elementy > 0:  ")
	fmt.Println(count2([]int{1,1,2,3,4,4,5,5,5}))

	fmt.Println("MinMax:  ")
	fmt.Println(minMax([]int{9,1,10,7,8}))

}

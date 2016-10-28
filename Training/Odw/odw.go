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

func findFirstX(tab []int, x int) (string,int){

	i := 0
	for i < len(tab){
		if tab[i] == x {
			return "Znaleziono x pod indeksem:", i
		}
		i++
	}
	return "Nie znaleziono x:",-1
}

func binarySearch(tab []int, x int) int{

	lower := 0
	upper := len(tab)-1


	for lower <= upper {
		middle := lower + (upper - lower)/2

		if tab[middle] == x {


			return middle

		}

		if tab[middle] > x {

			upper = middle - 1
		}

		if tab[middle] < x {
			lower = middle + 1
		}
	}

	return -1
}


func sumDigit(n int) int64 {

	var wynik int64

	str := strconv.Itoa(n)
	for _, i := range str{
		tmp, _:= strconv.ParseInt(string(i),10,64)
		wynik +=tmp
	}

	return wynik
}

func sumDigit2(n int) int {

	wynik :=0

	for n != 0{
		wynik += n%10
		n = n/10

	}

	return wynik
}


func pow(a float64, b float64) float64{

	wynik := 1.0

	if b > 0 {
		for b > 0 {
			wynik *= a
			b -=1
		}

	}

	if b < 0 {
		for b < 0 {
			wynik *= a
			b +=1
		}
		wynik = 1/wynik

	}

	return  wynik
}


func sr(n []float64) float64{

	var wynik float64
	var li float64

	for _,i := range n {
		wynik +=i
		li +=1

	}

	return wynik/li

}

func even(n []int) string{

	str := ""
	for _ , i := range n {
		if i % 2 == 0{
			str += strconv.Itoa(i)+ " jest parzysta \n"
		}

	}
	return  str
}

func dominanta(n []int)int{

	var l [10]int //wartosc
	var w [10]int //ilosc



	for i := 0; i<len(n) ; i++{

			if !search(l,n[i]){
				l[i] = n[i]
			}
	}

	for j:=0;j<len(l);j++ {

		for k:=0; k<len(n);k++{
			if l[j] == n[k]{
				w[j]+=1
		}
		}
	}

	max := w[0]
	in := 0
	for l:=1;l<len(w);l++{
		if w[l] > max{
			max = w[l]
			in = l
		}
	}

	return  l[in]
}

func search(n [10]int, x int) bool {

	s := false

	for i :=0;i<len(n);i++{
		if n[i] == x{
			s = true

		}

	}

	return s
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

	fmt.Println("Pierwszy element:  ")
	fmt.Println(findFirstX([]int{9,7,10,7,8},11))
	fmt.Println(findFirstX([]int{9,7,10,7,8},10))


	fmt.Println("Wyszukiwanie:  ")
	fmt.Println(binarySearch([]int{1,2,3,4,5,6,7,9,11},11))
	fmt.Println(binarySearch([]int{1,2,3,4,5,6,7},6))

	fmt.Println("Suma Cyfr: ")
	fmt.Println(sumDigit(148))
	fmt.Println(sumDigit2(1234))

	fmt.Println("Potęga: ")
	fmt.Println(pow(6,3))
	fmt.Println(pow(2,-1))

	fmt.Println("Srednia: ")
	fmt.Println(sr([]float64{1.2,2.5,3,4}))

	fmt.Println("Parzysta: ")
	fmt.Println(even([]int{1,2,3,4,5,6,7,9,11,8}))

	fmt.Println("Dominanta: ")
	fmt.Println(dominanta([]int{1,2,1,2,3,2}))

}

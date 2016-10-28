package main

import "fmt"

func odw(c string) string{
	r := []rune(c)

	for i , j := 0, len(r)-1 ; i < j; i,j = i+1,j-1  {
		r[i],r[j] = r[j], r[i]


	}

	return string(r)
}

func odw2(c string) string{
	r := []rune(c)

	for i := 0; i < len(r)/2 ; i++ {
		tmp := r[i]
		r[i] = r[len(r) -1 -i]
		r[len(r) -1 -i] = tmp

	}

	return string(r)
}

func bez(n int) int{

	if n > 0{
		return n
	}else if (n < 0) {

		return n*(-1)
	}
	return 0
}

func sum(s []float64) float64{

	var wynik float64 = 0

	for i := 0; i < len(s) ; i++ {
		wynik += s[i]

	}


	return wynik

}

func  sum2(s []float64) float64  {

	var wynik float64= 0.0
	for _,i := range s{
		wynik += i

	}
	return wynik

}

func main()  {

	fmt.Println(odw("abcd"),"-> dcba")
	fmt.Println(odw("abcde"),"-> edcba")
	fmt.Println(odw2("abcd"),"-> dcba")
	fmt.Println(odw2("abcde"),"-> edcba")


	fmt.Println(bez(-1))
	fmt.Println(bez(-(-1)))
	fmt.Println(bez(8))
	fmt.Println(bez(-8))


	fmt.Println(sum([]float64{1,2,3,-1,45}))
	fmt.Println(sum2([]float64{1,2,3,-1,45}))
}



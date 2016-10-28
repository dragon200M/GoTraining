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

	if n >= 0{
		return n
	}else if (n < 0) {

		return n*(-1)
	}
	return 0
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

}



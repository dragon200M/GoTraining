package main

import "fmt"

func main() {

	fmt.Println(odwInt(1234))
	fmt.Println(odwInt(123))
	fmt.Println(odwInt(1230))

	fmt.Println(odwNapis("abcde"))

	fmt.Println(odwNapis2("abcd"))

	fmt.Println(odwInt2(1234))
	fmt.Println(odwInt2(123))
	fmt.Println(odwInt2(1230))

}

func odwInt(n int) int {
	tmp := n
	wy := 0
	r := 0
	i := 0
	p := 1

	for tmp > 0 {
		tmp = tmp / 10
		i++
	}

	for i > 0 {
		j := i - 1

		for j > 0 {
			p = p * 10
			j--
		}

		r = n % 10
		wy += p * r
		n = n / 10
		p = 1
		i--

	}

	return wy
}

func odwNapis(txt string) string {
	r := []rune(txt)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[i]
		r[i] = r[len(r)-1-i]
		r[len(r)-1-i] = tmp

	}

	return string(r)
}


func odwNapis2(txt string) string {

	r :=[]rune(txt)

	for i,j:=0,len(r)-1;i<j;i,j=i+1,j-1{
		r[i],r[j]=r[j],r[i]

	}
	return string(r)

}


func odwInt2(n int) int{
	tmp:=n
	i:=0


	for tmp > 0 {
		tmp=tmp/10
		i++
	}

	p := 1
	r := 0
	w := 0

	for i > 0 {
		j := i -1

		for j > 0{
			p=p*10
			j--
		}
		r = n%10
		n=n/10
		w+=p*r
		p = 1
		i--
	}
	return w

}
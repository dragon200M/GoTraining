package main

import "fmt"

type person struct{
	first string
	last string
	age int
	addr


}
type  addr struct {
	country string
	city
}

type city struct {
	city string
}

func main() {
	a := addr{"Poland",city{"Cracow"}}
	me := person{"Maciej","Michalik",27,addr{"poland",city{"Cracow"}}}
	me2 := person{"Maciej","Michalik",27,a}
	fmt.Println(me2)
	fmt.Println(me)
	fmt.Println(me.first,me.last,me.age,me.country,me.city.city)


}

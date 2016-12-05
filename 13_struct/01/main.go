package main

import "fmt"

type person struct{
	first string
	last string
	age int
	adres addr
}
type  addr struct {
	country string
}

func main() {
	a := addr{"Poland"}
	me := person{"Maciej","Michalik",27,a}
	fmt.Println(me)
	fmt.Println(me.first,me.last,me.age,me.adres.country)


}

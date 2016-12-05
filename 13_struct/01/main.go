package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func main() {
	me := person{"Maciej","Michalik",27}
	fmt.Println(me)
	fmt.Println(me.first,me.last,me.age)


}

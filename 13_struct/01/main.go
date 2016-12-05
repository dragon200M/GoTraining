package main



import (
	"fmt")


type person struct{
	first string
	last string
	age int
	addr


}

func (p *person) sayHi() {
	fmt.Println("hi, my name is",p.first)
}




type  addr struct {
	country string
	city
}

type city struct {
	city string
}

type postCode interface {

}

func main() {
	a := addr{"Poland",city{"Cracow"}}
	me := person{"Maciej","Michalik",27,addr{"poland",city{"Cracow"}}}
	me2 := person{"Maciej","Michalik",27,a}
	fmt.Println(me2)
	fmt.Println(me)
	fmt.Println(me.first,me.last,me.age,me.country,me.city.city)

	me.sayHi()


}

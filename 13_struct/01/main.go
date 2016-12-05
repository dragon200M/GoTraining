package main



import (
	"fmt")


type person struct{
	first string
	last string
	age int
	addr


}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) sayHi() {
	fmt.Println("hi, my name is",p.first)
}

func (d doubleZero) sayHi()  {
	fmt.Println("hi",d.LicenseToKill)
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




	p1:=person{
		first :"M",
		last:"M",
		age: 27,
		addr: addr{
			country:"Poland",
			city: city{
				city:"Wieliczka",
			},
		},


	}


	fmt.Println(p1)

	p2:=doubleZero{
		person: person{
		first :"M",
		last:"M",
		age: 27,
		addr: addr{
			country:"Poland",
			city: city{
				city:"Wieliczka",
			},
		},
		},
	LicenseToKill: true,
	}

	fmt.Println(p2)

	p1.sayHi()
	p2.sayHi()
	p2.person.sayHi()
}

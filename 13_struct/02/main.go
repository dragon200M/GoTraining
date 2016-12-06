package main

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
)


type Person struct {
	First string
	Last string
	Age int
	//notExported int //ponieważ z małej litery nie jest eksportowane, Person jest objektem
	A int `json:"score"` //"-" pomijamy, "string" nazwa pola
}

type Animal struct {
	Name string
	Name2 string
}

type Dog struct {
	Animal
	Voice string


}


func main() {
	//Marshal
	p1:=Person{"Maciek","Michalik",27,6}
	bs,_:=json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T \n",bs)
	fmt.Println(string(bs))


	//UnMarshal
	var p2 Person
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
	//fmt.Println(p2.notExported)
	//fmt.Println(p2.A)



	bs2 :=[]byte(`{"First":"Maciek","Last":"Michalik","Age":27,"score":7}`)
	json.Unmarshal(bs2, &p2)

	fmt.Println("unmarshal")
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
	//fmt.Println(p2.notExported)
	fmt.Println(p2.A)

	fmt.Println("-----------------")
	p3 := Person{"Maciek","Michalik",27,6}
	json.NewEncoder(os.Stdout).Encode(p3)

	var p4 Person
	bs3 :=strings.NewReader(`{"First":"Maciek","Last":"Michalik","Age":27,"score":7}`)
	json.NewDecoder(bs3).Decode(&p4)

	fmt.Println(p4.First)
	fmt.Println(p4.Last)
	fmt.Println(p4.Age)
	fmt.Println(p4.A)

	fmt.Println("Dog")
	d1 := Dog{
		Animal: Animal{"Rocky",
		},Voice:"Wrrr",

	}

	json.NewEncoder(os.Stdout).Encode(d1)

}

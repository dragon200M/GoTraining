package main

import (
	"fmt"
	"encoding/json"
)


type Person struct {
	First string
	Last string
	Age int
	notExported int //ponieważ z małej litery nie jest eksportowane, Person jest objektem
	A int `json:"score"` //"-" pomijamy, "string" nazwa pola
}

func main() {

	p1:=Person{"Maciek","Michalik",27,007,5}
	bs,_:=json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T \n",bs)
	fmt.Println(string(bs))

}

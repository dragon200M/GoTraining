package main

import "fmt"

func main() {

	m:=make(map[string]int)

	m["k1"] = 1
	m["k1"] = 2
	m["k2"] = 2
	delete(m,"k1")

	v,k :=m["k2"]
	fmt.Println(v,k)



	m2 :=map[string]int{"foo":1,"bar":2}

	v1,k1:=m2["foo"]
	fmt.Println(v1,k1)
}
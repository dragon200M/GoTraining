package main

import (
	"fmt"
	"sort"
)

type People struct {
	name string
}

type ByName []People

func(a ByName) Len() int {return  len(a)}
func(a ByName) Swap(i,j int){a[i], a[j] = a[j], a[i] }
func(a ByName) Less(i,j int)bool{ return a[i].name < a[j].name}

func main() {

	studyGroup := []People{{"Zeno"},{"John"},{"All"},{"Jenny"}}


	fmt.Println(studyGroup)
	sort.Sort(ByName(studyGroup))
	fmt.Println(studyGroup)


}

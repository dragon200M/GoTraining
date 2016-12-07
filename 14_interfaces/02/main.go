package main

import (
	"fmt"
	"sort"
)

type People struct {
	name string
}

type ByName []People

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].name < a[j].name }

type ByNameS []string

func (a ByNameS) Len() int           { return len(a) }
func (a ByNameS) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNameS) Less(i, j int) bool { return a[i] < a[j] }

type IntSlice []int

func (a IntSlice) Len() int           { return len(a) }
func (a IntSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntSlice) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	studyGroup := []People{{"Zeno"}, {"John"}, {"All"}, {"Jenny"}}

	fmt.Println(studyGroup)
	sort.Sort(ByName(studyGroup))
	fmt.Println(studyGroup)

	studyGroup2 := []string{"Zeno", "John", "All", "Jenny"}

	sort.Sort(ByNameS(studyGroup2))
	fmt.Println(studyGroup2)

	studyGroup3 := []string{"Zeno", "John", "All", "Jenny"}
	sort.Strings(studyGroup3)
	fmt.Println(studyGroup3)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	sort.Sort(IntSlice(n))

	fmt.Println(n)

	n2 := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	sort.Ints(n2)
	fmt.Println(n2)

}

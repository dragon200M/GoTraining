package main

import "fmt"


func main() {

	t :=make([]int,0,5)

	fmt.Println(len(t),cap(t))

	for i:=0;i<80;i++{
		t = append(t,i)
		fmt.Println("Len:",len(t),"Cap:",cap(t),"Value:",t[i])
	}

	t2 := []string{"a","b","c","d","e"}

	fmt.Println(t2[0:2])
	fmt.Println(t2[2:])



}

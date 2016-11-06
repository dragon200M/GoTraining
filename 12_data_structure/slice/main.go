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


	record :=make([][]string,0)

	s1 :=make([]string,4)
	s1[0]= "a"
	s1[1]= "a"
	s1[2]= "a"
	s1[3]= "a"

	s2 :=make([]string,4)
	s2[0]= "b"
	s2[1]= "b"
	s2[2]= "b"
	s2[3]= "b"

	record = append(record,s1)
	record = append(record,s2)

	fmt.Println(record)



}

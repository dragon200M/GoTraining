package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 12
	var y  = 23.3131


	fmt.Println(y+float64(x))
	fmt.Println(int(y)+x)
	//conversion

	i,err :=strconv.Atoi("56")

	if(err != nil){
		fmt.Println(err)
	}

	fmt.Println(i)

	fmt.Println([]byte("56"))


}

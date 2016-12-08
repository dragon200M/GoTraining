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

	fmt.Printf("%T, %s \n",string([]byte{53,54}),string([]byte{53,54}))
	//fmt.Println(int([]byte{53,54}))//bład
	fmt.Println(strconv.Atoi(string([]byte{53,54})))

	i2,err2 :=strconv.Atoi(string([]byte{53,54}))

	if(err2 != nil){
		fmt.Println(err)
	}
	fmt.Printf("%T, %d \n",i2,i2)

}

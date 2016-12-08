package main

import (
	"fmt"
	"strconv"
)

func main() {
	//conversion
	var x = 12
	var y  = 23.3131


	fmt.Println(y+float64(x))
	fmt.Println(int(y)+x)

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

	//assertion

	var name interface{} = "Maciej"

	str, ok := name.(string)

	if ok {
		fmt.Printf("%T, %s \n",str,str)
	}else{
		fmt.Println("To nie jest string")
	}

	var n interface{} = 45

	str2, ok2 := n.(int)

	if ok2 {
		fmt.Printf("%T, %d \n",str2,str2)
	}else{
		fmt.Println("To nie jest int")
	}

}

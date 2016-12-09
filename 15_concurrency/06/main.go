package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {

	c :=make(chan  int)
	s :=make(chan string)
	var zm int

	go func(){
		for i:=0;i<10;i++{
			c<-i

			zm = i*i*i
			s<-strconv.Itoa(zm)

		}
	}()

	go func(){
		for {
			fmt.Println(<-c,<-s)

		}
	}()


	time.Sleep(3*time.Second)


}

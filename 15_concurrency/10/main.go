package main

import "fmt"



func main() {
	c :=make(chan int)
	c1 :=make(chan int)

	go func(){
		c <- 1
	}()


	fmt.Println(<-c)





	go func(){
		for i:=0;i<10;i++ {
			c1 <- i

		}
	}()


	for{
		fmt.Println(<-c1)
	}
}

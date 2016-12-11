package main

import "fmt"

func main() {

	c1 := make(chan int)

	go func(){
		for i:=0;i<10;i++ {
			c1 <- i

		}
		close(c1)
	}()



	for a := range c1 {

		fmt.Println( a)
	}

}

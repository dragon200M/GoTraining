package main

import "fmt"

func main() {
	c:=make(chan int)
	done :=make(chan bool)


	for j:=0;j<10;j++{
	go func() {
		for i := 0; i < 10; i++ {
			c<-i
		}
		done <-true
	}()

	}

	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()

	for d :=range c{
		fmt.Println(d)
	}

}

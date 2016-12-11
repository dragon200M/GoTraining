package main

import (
	"fmt"
)

func main() {

	c1 := inc("Foo:")
	c2 := inc("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)

	fmt.Println(<-c3 , <-c4)


}
func puller(in chan int) chan int {
	out := make(chan int)
	go func() {
		var suma int
		for s := range in {
			suma += s

		}

		out <- suma

		close(out)
	}()
	return out

}
func inc(s string) chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			fmt.Println(s, i)

		}

		close(out)

	}()
	return out
}

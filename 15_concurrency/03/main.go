package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
	"math/rand"
)

var wg sync.WaitGroup
var counter int

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())

}

func main() {
	wg.Add(2)
		go inc("foo")
		go inc("bar")
	wg.Wait()
	fmt.Println("counter:",counter)
}



func inc(s string){
	for i:=0;i<10;i++ {
		x:=counter
		x++

		time.Sleep(time.Duration(rand.Intn(30))*time.Millisecond)
		counter=x
		fmt.Println(s,i,"counter:",counter)
	}
	wg.Done()
}


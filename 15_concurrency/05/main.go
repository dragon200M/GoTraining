package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64




func main() {
	wg.Add(2)
		go inc("foo")
		go inc("bar")
	wg.Wait()
	fmt.Println("counter:",counter)
}



func inc(s string){
	for i:=0;i<10;i++ {
		time.Sleep(time.Duration(rand.Intn(30))*time.Millisecond)

		atomic.AddInt64(&counter,1)
		fmt.Println(s,i,"counter:",counter)

	}
	wg.Done()
}


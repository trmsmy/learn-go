//main function

package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var w = sync.WaitGroup{}

func main20() {
	//iterate 10 times
	for i := 0; i < 10; i++ {
		w.Add(2)
		go op1()
		go op2()
		w.Wait()
	}
}

func op1() {

	defer w.Done()
	//sleep 2 seconds
	time.Sleep(2 * time.Second)
	m.Lock()
	fmt.Println("op1")
	m.Unlock()
}
func op2() {
	defer w.Done()

	//sleep 2 seconds
	time.Sleep(2 * time.Second)
	m.Lock()
	fmt.Println("op2")
	m.Unlock()
}

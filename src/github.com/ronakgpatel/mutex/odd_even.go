package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.Mutex{}
var counter1 = 1
var counter2 = 0

func main() {
	//runtime.GOMAXPROCS(2)
	//fmt.Printf("Threads : # %v\n", runtime.GOMAXPROCS(-1))

	for counter2 < 10 {
		wg.Add(2)
		m.Lock()
		go printEven()
		m.Lock()
		go printOdd()
	}
	wg.Wait()
}

func printEven() {
	fmt.Printf("%v\n", counter2)
	counter2 += 2
	m.Unlock()
	wg.Done()
}

func printOdd() {
	fmt.Printf("%v\n", counter1)
	counter1 += 2
	m.Unlock()
	wg.Done()
}

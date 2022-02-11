package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
	wg.Add(1)
	var msg = "HELLO"
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()

	//time.Sleep(100 * time.Millisecond)
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello")
}

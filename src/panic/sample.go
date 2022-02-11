package main

import (
	"fmt"
	"log"
)

func main() {
	/*a, b := 1, 0
	fmt.Println(a / b)*/

	/*fmt.Println("Pre processing")
	defer fmt.Println("This should be executed last")
	panic("Oh something bad happened")

	fmt.Println("Post processing")*/

	fmt.Println("Beginning")
	panicker()
	fmt.Println("Ending")

}

func panicker() {

	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Got Error:", err)
			//panic(err)
		}
	}()

	panic("Some error occured")
	fmt.Println("Done panicking")

}

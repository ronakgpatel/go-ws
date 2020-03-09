package main

import "fmt"

func main() {
	defer fmt.Println("First")
	fmt.Println("Second")
	sample_func()

	defer fmt.Println("Third")

	//the below would output "Start" and ****NOT**** "end"
	a := "Start"
	defer fmt.Println(a)
	a = "end"

}

func sample_func() {
	fmt.Println("Inside sample_func")
	defer fmt.Println("Exiting sample_func")
}

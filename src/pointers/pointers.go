package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println("Address : ", &a, b)
	fmt.Println("Values :", a, *b)
	a = 18
	fmt.Println("Values :", a, *b)
	*b = 22
	fmt.Println("Values :", a, *b)

	var arr [3]int = [3]int{1, 2, 3}
	arrB := &arr[0]
	arrC := &arr[1]
	fmt.Printf("%v %p %p\n", arr, arrB, arrC)
}

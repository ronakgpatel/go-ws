package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6} // slice
	b := a                       //by default shallow copy, opposite of array
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length:%v\n", len(a))
	fmt.Printf("Capacity:%v\n", cap(a))

	b1 := a[:]  //slice of all elements
	c := a[3:]  //slice from 4th element to end
	d := a[:6]  //slice first 6 elements
	e := a[3:6] //slice the 4,5,6 th elements

	fmt.Println(b1)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("-----------------")
	//make function
	a1 := make([]int, 3, 100)
	fmt.Println(a1)
	fmt.Printf("Length : %v\n", len(a1))
	fmt.Printf("Capacity: %v \n", cap(a1))
	a1 = append(a1, 39, 98, 99)
	fmt.Println(a1)
	fmt.Printf("Length : %v\n", len(a1))
	fmt.Printf("Capacity: %v \n", cap(a1))

	a1 = append(a1, []int{100, 200, 300}...)
	fmt.Println(a1)
	fmt.Printf("Length : %v\n", len(a1))
	fmt.Printf("Capacity: %v \n", cap(a1))

	arr := [5]int{1, 2, 3, 4, 5}
	brr := append(arr[:3], arr[4:]...)
	fmt.Println(brr)
	fmt.Println(arr)
}

package main

import "fmt"

func main() {
	grades := [3]int{97, 92, 85} // number of elements & type
	fmt.Printf("Grades: %v \n", grades)

	moregrades := [...]int{10, 30, 97, 92, 85} // number of elements & type
	fmt.Printf("More  Grades: %v \n", moregrades)

	var students [100]string
	fmt.Printf("Students:%v\n", students)

	students[0] = "Lisa"
	students[2] = "Murugan"
	students[1] = "Subbu"
	fmt.Printf("Students : %v\n", students)
	fmt.Printf("First Student : %v\n", students[1])
	fmt.Printf("Total Students : %v\n", len(students))

	//multi diemnsion array
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}
	fmt.Printf("Identity Matrix : %v\n", identityMatrix)

	a := [...]int{1, 2, 3}
	b := a //deep copy

	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)

	c := &a // shallow copy
	c[1] = 50
	fmt.Println(a)
	fmt.Println(c)
}

package main

import "fmt"

func main() {
	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
	(*ms).foo = 42
	fmt.Println(ms)
	fmt.Println("Actual Value :", (*ms).foo)

	ms.foo = 24 // syntactic sugar, the imact is same as in line 10
	fmt.Println(ms)
	fmt.Println("Actual Value :", ms.foo)

}

type myStruct struct {
	foo int
}

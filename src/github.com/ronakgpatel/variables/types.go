package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)
	var m bool
	fmt.Printf("%v, %T\n", m, m)
	var n4 int8 = 0
	var n1 int16 = 0
	var n2 int32 = 0
	var n3 int64 = 0
	fmt.Printf("%v, %T\n", n4, n4)
	fmt.Printf("%v, %T\n", n1, n1)
	fmt.Printf("%v, %T\n", n2, n2)
	fmt.Printf("%v, %T\n", n3, n3)
	var n5 uint8 = 100 //unsinged type, unit8,unit16,unit32,unit64
	fmt.Printf("%v, %T\n", n5, n5)

	a := 10
	b := 120
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	var a1 int = 10
	var a2 int8 = 3
	fmt.Println(a1 + int(a2))

	var s string = "Hi this is first string"
	fmt.Printf("%v, %T\n", string(s[0]), s[0])

	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}

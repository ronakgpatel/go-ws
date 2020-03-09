package main

import "fmt"

func main() {
	var c complex64 = 1 + 2i + 45i

	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	var c1 complex128 = 102.5i

	fmt.Printf("%v, %T\n", c1, c1)
	fmt.Printf("%v, %T\n", real(c1), real(c1))
	fmt.Printf("%v, %T\n", imag(c1), imag(c1))

	var c2 complex64 = complex(14, 5)
	fmt.Printf("%v, %T\n", c2, c2)

}

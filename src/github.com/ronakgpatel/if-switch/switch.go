package main

import "fmt"

func main() {
	//break keyword not required as its implicit
	//break keyword can be used to break within single case block
	switch i := 32 / 8; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 0, 2, 4, 6, 8:
		fmt.Println("Even")
	default:
		fmt.Println("can not evaluate")
	}

	j := 3
	switch {
	case j%2 == 0:
		fmt.Println("j is even")
	case j%2 == 1:
		fmt.Println("j is odd")
		fallthrough
	case j < 0:
		fmt.Println("***j does not matter***")
	default:
		fmt.Println("IDK - I don't know")
	}
}

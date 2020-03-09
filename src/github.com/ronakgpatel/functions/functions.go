package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	ans, err := add(1, 3)
	fmt.Println("Answer:", ans, "Error:", err)
	ans, err = add(0, 3)
	fmt.Println("Answer:", ans, "Error:", err)

	g := greeting{
		name: "Briti",
		msg:  "Hello",
	}
	//
	g.greet()
	fmt.Println(g.msg, " ", g.name)
	fmt.Println("1.2+6.3 =", add_float(1.2, 6.3))
	print_strings("Hi", "how", "are", "you?")

	func() {
		fmt.Println("From anonymous function")
	}()
	my_func := func(a int, b string) {
		fmt.Println("First arg : ", a)
		fmt.Println("Second arg : ", b)
	}
	my_func(1, "One")

	//var f func(string,string,int) (int,error)

}

type greeting struct {
	name string
	msg  string
}

func (g *greeting) greet() {
	fmt.Println("Greeting", g.msg, " ", g.name)
	g.name = "Alexa"
}

func add(a int, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Both must be non-zero")
	}
	return a + b, nil
}

func add_float(a, b float64) float64 {
	return a + b
}

func print_strings(data ...string) {
	fmt.Println("Arguments :", data)
	fmt.Println("Total args:", len(data))
}

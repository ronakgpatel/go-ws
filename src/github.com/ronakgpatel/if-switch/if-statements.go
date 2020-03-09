package main

import "fmt"

func main() {
	if true {
		fmt.Println("The test is true")
	} else {
		fmt.Println("The test is false")
	}

	statePopulations := map[string]int{
		"A": 100,
		"B": 200,
	}
	if pop, ok := statePopulations["A"]; ok {
		fmt.Println(pop)
	}

}

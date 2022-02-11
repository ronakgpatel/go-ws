package main

import "fmt"

func main() {
	var i interface{} = [2]int{1, 2}
	switch i.(type) {
	case int:
		/*if i > 128 || i < -127 {
			fmt.Println(" i is int but its too big")
		}
		break*/
		fmt.Println(" i is int")
	case float64:
		fmt.Println(" i is float64")
	case string:
		fmt.Println(" i is string")
	case [3]int:
		fmt.Println(" i is [3]int")
	default:
		fmt.Println(" i is not recognizable")
	}
}

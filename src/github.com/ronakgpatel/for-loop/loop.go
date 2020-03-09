package main

import "fmt"

func main() {
	for i, j, k := 0, 0, 0; i < 5; i, j, k = i+1, j+2, k+3 {
		fmt.Println(i, j, k)
	}

	z := 10
	for {
		if z < 0 {
			break
		} else if z < 10 {
			fmt.Println(z)
			z = z - 1
		} else {
			z = z - 1
		}
	}
	fmt.Println("******")
Loop:
	for {
		for k := 0; k < 10; k++ {
			if k == 2 {
				break Loop
			}
			fmt.Println(k)

		}
	}

	fmt.Println("******")

	s := []int{10, 20, 30, 40} // this can be slice, map or array
	for k, v := range s {
		fmt.Println(k, v)
	}

	pops := map[string]string{
		"Delhi":     "4Cr",
		"Mathura":   "10L",
		"Rishikesh": "5L",
	}
	fmt.Println("******")
	for k, v := range pops {
		fmt.Println(k, v)
	}
}

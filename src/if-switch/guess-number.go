package main

import "fmt"

func main() {

	// &&  - is same as in java , logical and
	// || - logical or
	// ! - Not
	// short circuting is done by Go as well.

	number := 50
	guess := -10

	if guess < 1 {
		fmt.Println("Invalid guess, must be higher than 1")
	} else if guess > 100 {
		fmt.Println("Invalid guess, must be less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}

		if guess == number {
			fmt.Println("You got it")
		}
	}

	fmt.Println(guess < number, guess > number, guess == number)
}

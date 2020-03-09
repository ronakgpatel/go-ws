package main

import "fmt"

func main() {

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Delhi":   35000000,
		"Mumbai":  60000000,
		"Amdavad": 4500000,
		"Shilong": 450000,
	}
	statePopulations["Manali"] = 1000
	copyOfData := statePopulations
	fmt.Println("Original: ", statePopulations)
	fmt.Println("Copy(before delete) : ", copyOfData)

	delete(statePopulations, "Manali")

	fmt.Println("Copy(after delete): ", copyOfData)
	fmt.Println(statePopulations["Manali"])
	manaliPopulation, ok := statePopulations["Manali"]
	fmt.Println(manaliPopulation, ok)
	fmt.Println(len(statePopulations))

}

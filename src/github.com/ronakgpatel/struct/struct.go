package main

import "fmt"

type Movie struct {
	year       int
	budget     float64
	actorNames []string
}

func main() {

	aMovie := Movie{
		year:   2019,
		budget: 50000000.0,
		actorNames: []string{
			"Vicky",
			"Yamini",
			"Paresh",
		},
	}
	fmt.Println(aMovie)
	fmt.Println("Release Year : ", aMovie.year)
	fmt.Println("One of the actor : ", aMovie.actorNames[1])

	//anonymous struct
	newMovie := struct {
		year int
		name string
	}{year: 1999, name: "Ashok Kumar"}
	fmt.Println(newMovie)
	anotherMovie := newMovie

	newMovie.name = "Anil"
	fmt.Println("New ", newMovie)
	fmt.Println("Copied ", anotherMovie)
}

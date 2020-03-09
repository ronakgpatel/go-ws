package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name    string `required max:"100"` /// tag something like Annotation
	Origian string
}

type Bird struct {
	Animal   //this is just specifying the embedding of the type into this type
	CanFly   bool
	SpeedKPH float32
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origian = "Australia"
	b.CanFly = false
	b.SpeedKPH = 48.0
	fmt.Println(b)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	b1 := Bird{
		Animal:   Animal{Name: "Emu", Origian: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b1.Name)
}

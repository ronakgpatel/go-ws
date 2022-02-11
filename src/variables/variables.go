package main

import (
	"fmt"
	"strconv"
)

//package level variable
//starting with UpperCase - Public Access
//starting with lowercase - Accessible with in package
//no prive scope
//package level variable
var sessionVariable int = 10
var actorName string = "I don't know"

var (
	blockedVairables int = 100
	director             = "TBD"
	reading              = 100.94
)

func main() {
	//declare variables - block scope (local vairables)
	var local int
	var local1 int = 42
	local = 99
	local2 := 34.2
	fmt.Printf("%v, %T\n", local, local)
	fmt.Printf("%v, %T\n", local1, local1)
	fmt.Printf("%v, %T\n", local2, local2)

	//convert variables
	var i int = 42
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)

	//String conversion
	var s string = string(i)
	fmt.Printf("%v, %T \n", s, s)
	s = strconv.Itoa(i)
	fmt.Printf("%v (as a string), %T \n", s, s)
}

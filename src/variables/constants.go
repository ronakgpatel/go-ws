package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
	d = iota
	e
	f
)
const (
	a1 = iota
	b1
	c1
)

const (
	_ = iota + 10 //blank identifier variable which can be thrown away
	catSpecialists
	dogSpecialists
	snakeSpecialists
)

const (
	_  = iota
	KB = 1 << (10 * iota) // shift 1 bit ie. mulitply by 2(2^1)
	MB                    // shift 2 bits ie. mulitply by 4(2^2)
	GB                    // shift 3 bits ie. mulitply by 8(2^3)
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadQuarters
	canSeeFinancials
)

func main() {
	const a int = 10
	const b float64 = 100.048
	const c bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("d=%v\n", d)
	fmt.Printf("f=%v\n", f)

	fmt.Printf("a1=%v\n", a1)

	fmt.Printf("snakeSpecialists=%v\n", snakeSpecialists)

	fileSize := 4000000000. //remove the dot and check the output as well
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials
	fmt.Printf("Roles : %b\n", roles)
	fmt.Printf("IsAdmin : %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("isHeadQuarters : %v\n", isHeadQuarters&roles == isHeadQuarters)
}

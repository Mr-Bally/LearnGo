package main

import (
	"fmt"
)

const (
	// Big number = 1 bit shifted left 100 places
	Big = 1 << 100
	// Small numver = Shift Big right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

/*
	Const are declared using the "const" keyword.
	The constants cannot be declared using ":=" syntax
*/
func main() {
	const Truth = true
	const Lang = "GO"

	fmt.Println("Signed constants")
	fmt.Println(Truth, Lang)

	/*
		Notice we do not try to call needInt on const Big
		as this would cause an overflow
	*/
	fmt.Println("Unsigned constants")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/*
	Unsigned const types take the value of the context they
	are needed
*/
func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

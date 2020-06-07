package main

/*
	You can factor imports into a single statement,
	rather than several "imports" statements
*/
import (
	"fmt"
	"math"
)

/*
	A name is exported if it begins with a capital letter.
	In this example the math package exports Pi
*/
func main() {
	fmt.Println(math.Pi)
}

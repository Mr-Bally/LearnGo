package main

import (
	"fmt"
)

/*
	Go with try and infer the type from
	the right hand side of an assignment
*/
func main() {
	i := 42              // int
	f := 3.142           // float64
	g := 0.867 + 0.5i    // complex128
	s := "StringExample" // string
	fmt.Println(i, f, g, s)
}

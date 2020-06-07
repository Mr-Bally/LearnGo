package main

import "fmt"

/*
	Here are the most common types you will use
	in Go. There are more (e.g. "uint8") but for now
	we will just examine the default values for these
*/
func main() {
	var i int
	var f float64
	var b bool
	var s string

	/*
		Notice we are using the print format function
		with formatting argument ("%v" means default
		value representation in this case)
	*/
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

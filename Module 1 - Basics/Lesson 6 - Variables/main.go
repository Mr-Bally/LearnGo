package main

import "fmt"

/*
	Variables are declared using "var".
	This creates a list of variables with
	the type coming last. This can be done
	at the package or function level
*/
var c, python, java bool

func main() {
	var i int
	fmt.Println("Package and function variables")
	fmt.Println(c, python, java, i)

	fmt.Println("\nSame type variable")
	fmt.Println(intialiseSameType())

	fmt.Println("\nMixed type variable")
	fmt.Println(intialiseDifferentType())

	fmt.Println("\nImplicit variables")
	fmt.Println(getImplicitType())
}

/*
	Here we set variables of the same type.
	You can omit the type as the variable will
	take the place of the initialiser
*/
func intialiseSameType() (int, int) {
	var intOne, intTwo = 1, 2
	return intOne, intTwo
}

/*
	You can also do this with variables of
	different types
*/
func intialiseDifferentType() (bool, int, string) {
	var a, b, c = true, 5, "ExampleString"
	return a, b, c
}

/*
	You can declare a variable using "var" keyword
	or use the implicit ":=" method. This can only
	be done within a function
*/
func getImplicitType() (string, int) {
	k := 3
	testString := "Implicit string"

	return testString, k
}

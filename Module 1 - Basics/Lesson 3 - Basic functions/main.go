package main

import "fmt"

/*
	This function takes two integers as arguments.
	Notice the type comes before the variable name
	and that the function return type is also after
	the parameter list
*/
func add(x int, y int) int {
	return x + y
}

/*
	You can shorten the argument list by grouping the same types together
*/
func betterAdd(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Add function", add(5, 6))
	fmt.Println("Better add function", betterAdd(5, 6))
}

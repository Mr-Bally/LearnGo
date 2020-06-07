package main

import "fmt"

/*
	This is a basic for loop, the loop stops
	when the bool condition evaluates to false.
	Notice the lack of parentheses like most
	other languages
*/
func main() {
	fmt.Println("Basic for loop")
	fmt.Println(basicLoop())

	fmt.Println("Implicit for loop")
	fmt.Println(implicitForLoop())
}

func basicLoop() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func implicitForLoop() int {
	sum := 1
	/*
		In this example the loop continues until
		sum is greater than 1000
	*/
	for sum < 1000 {
		sum += sum
	}
	return sum
}

/*
	Without any conditions this will loop
	forever
*/
func infinite() {
	for {
	}
}

package main

import "fmt"

/*
	You can define multiple result types for your
	returning function
*/
func swap (a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("WordOne", "WordTwo")
	fmt.Println(a, b)
}

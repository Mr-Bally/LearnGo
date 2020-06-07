package main

import "fmt"

/*
	Notice here the return statement does not explictly
	return a result. This is because of the named variables
	at the top of the function. This should only be used
	in shorter function for readability
*/
func split(num int) (x, y int) {
	x = num * 1 / 4
	y = num - x
	return
}

func main() {
	fmt.Println("Split number 18")
	fmt.Println(split(18))
}

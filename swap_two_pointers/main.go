package main

import "fmt"

/*
Write a program to demonstrate pointer declaration and dereferencing.
Write a program to swap two numbers using pointers.
*/
func main() {
	var variable1 *int
	data1 := 5
	variable1 = &data1
	fmt.Println(*variable1)

	var variable2 *int
	data2 := 7
	variable2 = &data2
	fmt.Println(*variable2)

	//swap
	variable1, variable2 = variable2, variable1
	fmt.Println(*variable1, *variable2)
}

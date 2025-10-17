package main

import "fmt"

func main() {
	/*
		Write a program to show nil pointer dereference handling.
	*/
	var nil_pointer *int
	fmt.Println(*nil_pointer)
}

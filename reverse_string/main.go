package main

import "fmt"

func main() {
	//Write a program to reverse a string.
	/*
		Strings in Go are immutable. You cannot assign to stringdata[i] directly.
		string(stringdata[i]) converts a single byte to a string, but you cannot assign to it.
		To swap characters in a string, you need to convert the string to a slice of runes first.
	*/
	stringdata := "string"
	runes := []rune(stringdata)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	reversed := string(runes)
	fmt.Println("Original:", stringdata)
	fmt.Println("Reversed:", reversed)
}

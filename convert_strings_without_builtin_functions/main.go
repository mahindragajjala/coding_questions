package main

import "fmt"

func main() {
	/*
		Convert a string to uppercase without using the built-in function.
	*/
	/*
		'a' to 'z' → ASCII values from 97 to 122
		'A' to 'Z' → ASCII values from 65 to 90
	*/
	inputdata := "golang"
	runeinputdata := []rune(inputdata)
	fmt.Println(runeinputdata)
	newinputdata := []rune{}
	for i := 0; i < len(runeinputdata); i++ {
		newinputdata = append(newinputdata, runeinputdata[i]-32)
	}
	fmt.Println(newinputdata)
	fmt.Println(string(newinputdata))
}

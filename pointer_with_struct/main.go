package main

import "fmt"

//Create a pointer to a struct and modify its value.
type Structure_Value struct {
	Value1 int
	Value2 string
}

func main() {
	sv := Structure_Value{
		Value1: 2,
		Value2: "3",
	}
	pointer := &sv
	pointer.Value1 = 222
	pointer.Value2 = "333"

	fmt.Println(*pointer)
}

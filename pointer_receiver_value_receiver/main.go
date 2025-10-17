package main

import "fmt"

/*
Explain the difference between a pointer and value receiver with an example.
*/
func Pointer_Receiver() {
	var Receiver *string
	//we can store only address of the value only in the receiver variable.
	payload := "data"
	Receiver = &payload
	fmt.Println(Receiver)
}
func Value_Receiver() {
	var Value_Receiver string
	/*
		we can directly insert the data with any address.
	*/
	Value_Receiver = "data"
	fmt.Println(Value_Receiver)
}

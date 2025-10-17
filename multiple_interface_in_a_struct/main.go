package main

import "fmt"

type interface_Struct struct {
	interface_variable_one   interface{}
	interface_variable_two   interface{}
	interface_variable_three interface{}
}

func main() {
	//Show how to implement multiple interfaces on a struct.
	is := interface_Struct{
		interface_variable_one:   1,
		interface_variable_two:   "one",
		interface_variable_three: "o_n_e",
	}
	fmt.Println(is)
}

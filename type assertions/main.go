package main

import "fmt"

/*
i.(type) is not valid on its own in normal code.
It only appears inside a type switch, where Go allows you to inspect the dynamic type of an interface.
*/
func main() {
	var interface_data interface{} = "hello"
	switch v := interface_data.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("default", v)
	}

}

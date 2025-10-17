package main

import "fmt"

type structure struct {
	Value  int
	Value2 int
}

func main() {
	//Create a slice of structs and iterate over it.
	sliceOfData := []structure{{1, 2}, {3, 4}, {5, 6}}
	for i := 0; i < len(sliceOfData); i++ {
		fmt.Println(sliceOfData[i])
	}
}

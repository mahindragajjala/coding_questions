package main

import "fmt"

func Copy_array() {
	//Copy one array to another
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	new_array := [10]int{}

	array = new_array
	fmt.Println(array)
}

package main

import "fmt"

//Print all elements of an array
func main() {
	array := [5]int{1, 2, 3, 4, 5}

	//sequence
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i])
	}
	fmt.Println()
	fmt.Print("Even:")
	//even
	for i := 0; i < len(array); i++ {

		if array[i]%2 == 0 {
			fmt.Print(array[i])
		}
	}
	fmt.Println()
	fmt.Print("Odd:")
	//odd
	for i := 0; i < len(array); i++ {

		if array[i]%2 != 0 {
			fmt.Print(array[i])
		}
	}
}

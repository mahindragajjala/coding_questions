package main

import "fmt"

func main() {
	//Search for an element (Linear Search)
	numbers := []int{1, -2, 3, -4, 5, -6, 7, -8}
	seach_value := 5
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == seach_value {
			fmt.Println(i)
		}
	}
}

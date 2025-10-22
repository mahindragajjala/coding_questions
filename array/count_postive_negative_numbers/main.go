package main

import "fmt"

//Count positive and negative numbers
func main() {
	Count_postive_negative_numbers()
}
func Count_postive_negative_numbers() {
	numbers := []int{1, -2, 3, -4, 5, -6, 7, -8}
	var postive []int
	var negative []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			postive = append(postive, numbers[i])
		} else {
			negative = append(negative, numbers[i])
		}
	}
	fmt.Println(postive)
	fmt.Println(negative)
}

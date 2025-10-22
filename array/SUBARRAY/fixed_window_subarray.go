package main

import "fmt"

func Fixed_window_subarray(array [5]int) {
	fixed_value := 3
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			Sum_slice(array[i:j], fixed_value)
		}
	}
}
func Sum_slice(array []int, fixed_value int) {
	var sum int
	if len(array) == fixed_value {
		fmt.Println("array:", array)
		for i := 0; i < len(array); i++ {
			sum = sum + array[i]
		}
	}
	fmt.Println()
}

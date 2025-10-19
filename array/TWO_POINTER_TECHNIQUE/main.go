package main

import "fmt"

/*
The Two-Pointer Technique is a common approach in algorithms where two pointers (indexes)
traverse the array or string simultaneously to solve problems efficiently.

It’s especially useful in:
Sorted arrays
Subarray or substring problems

Problems involving sums, pairs, or duplicates
Types of Two-Pointer Approaches

a. Opposite-direction pointers
One pointer starts at the beginning, another at the end.
Move pointers toward each other based on a condition.
Use case: Find two numbers in a sorted array that sum to a target.

b. Same-direction (sliding) pointers
Both pointers start at the beginning.
One pointer moves faster or slower to maintain a condition (like a window size or sum).
Use case: Find subarrays with a sum ≤ target (Sliding Window).
*/
func Brute_force_approach(array []int) {
	window_size := 3
	var sum int
	for i := 0; i < len(array); i++ {
		sum = 0
		for j := i + 1; j <= len(array); j++ {
			window := array[i:j]
			if len(window) == window_size {
				fmt.Println(window)

				for i := 0; i < len(window); i++ {
					sum = sum + window[i]

				}
			}
		}
		fmt.Println(sum)
	}
}

/*
i=0, j=3 → [1 2 3]
i=1, j=4 → [2 3 4]
i=2, j=5 → [3 4 5]
*/
func Two_Pointer_Technique() {
	array := []int{1, 2, 3, 4, 5}
	windowSize := 3

	i := 0
	j := windowSize

	for j <= len(array) {
		fmt.Println(array[i:j])
		i++
		j++
	}
}

func Opposite_Direction() {
	arr := []int{1, 2, 3, 4, 5}
	i := 0
	j := len(arr) - 1

	for i < j {
		fmt.Printf("Pair: (%d, %d)\n", arr[i], arr[j])
		i++
		j--
	}
}
func Opposite_Direction_using_for_loop() {
	arr := []int{1, 2, 3, 4, 5}

	// Nested function to print a pair
	printPair := func(a, b int) {
		fmt.Printf("Pair: (%d, %d)\n", a, b)
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		printPair(arr[i], arr[j])
	}
}
func main() {
}

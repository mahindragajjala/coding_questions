package main

import (
	"fmt"
)

/*
The Sliding Window is a pattern/technique used to optimize loops over
arrays or strings when you need to find something in a contiguous
block of elements, like:

Maximum sum of k consecutive elements
Minimum length subarray with a given sum
Longest substring without repeating characters

Types of Sliding Windows
			"Fixed-Size Sliding Window"
			Window size is constant (k)
			Example: Maximum sum of k consecutive numbers
			Algorithm:
			Compute the sum of the first k elements.
			Move the window one step right:
			Add the new element
			Remove the first element of the previous window
			Repeat until the end of the array.
-------------------------------------------------------------------------------------------
			"Variable-Size Sliding Window"
			Window size can expand or shrink
			Useful when solving problems like:
			Minimum subarray length with sum ≥ target
			Longest substring with at most k distinct characters
			Algorithm (example: min length subarray sum ≥ target):
			Initialize two pointers: start and end
			Expand end pointer to include elements until condition is met
			Shrink start pointer to minimize the window while still satisfying condition
			Repeat until end reaches the array end


 Feature           Sliding Window                                                        Two Pointers
 Purpose           Used for contiguous subarrays or substrings (continuous block).       Used for searching or comparing elements — not necessarily contiguous.
 Window Nature     Focuses on a window (range) that expands or slides.                   Focuses on two independent pointers that move based on conditions.
 Usage             Finding sums, min/max subarray, longest substring, etc.               Finding pairs, intersections, merging sorted arrays, etc.
 Movement          Usually both pointers move forward only, maintaining a valid window.  One or both pointers may move independently in any pattern.
 Typical Problems  Maximum sum subarray, longest substring, min window substring.        Two-sum in sorted array, merging arrays, removing duplicates, etc.
 Contiguity        Always deals with contiguous ranges.                                  May deal with non-contiguous elements.

*/
// Maximum sum of k consecutive elements in an array.
func Maximum_sum_k_consecutive_elements_array(array []int) {
	window_size := 3
	i := 0
	j := window_size
	var Maximum int = 1
	for j <= len(array) {
		array := array[i:j]
		fmt.Println(array)
		var sum int
		for i := 0; i < len(array); i++ {
			sum = sum + array[i]
		}
		if sum > Maximum {
			Maximum = sum
		}
		i++
		j++
	}
	fmt.Println(Maximum)
}

// Minimum sum of k consecutive elements in an array.
func Minimum_sum_k_consecutive_elements_array(array []int) {
	window_size := 3
	i := 0
	j := window_size
	var Minimum int = 1
	for j <= len(array) {
		array := array[i:j]
		fmt.Println(array)
		var sum int
		for i := 0; i < len(array); i++ {
			sum = sum + array[i]
		}
		if sum < Minimum {
			Minimum = sum
		}
		i++
		j++
	}
	fmt.Println(Minimum)
}

// Count of subarrays of size k with sum equal to a target.
func Count_of_subarrays_of_k_with_sum_equal_to_a_target() {
	array := []int{1, 2, 3, 4, 5}
	window := 3
	i := 0
	j := window
	sum_of_subarrays := 5
	for j <= len(array) {
		var sum int
		for i := 0; i < j; i++ {
			sum = sum + array[i]

		}
		if sum > sum_of_subarrays {
			fmt.Println(array[i:j])
		}
		i++
		j++
	}
}

// First negative number in every window of size k.
func Negative_Number() {
	array := []int{1, -6, 2, -3, 3, -4, 5, -7, 6, -1, 7}
	window := 3
	i := 0
	j := window
	for j <= len(array) {
		value, index := find_the_negative_number(array[i:j])
		if value != 0 {
			fmt.Println(value, index)
		}
		i++
		j++
	}

}
func find_the_negative_number(array []int) (int, int) {
	for i := 0; i < len(array); i++ {
		if array[i] < 0 {
			return array[i], i
		}
	}
	return 0, 0
}

// Number of zeros in all subarrays of size k.
func Number_of_zeros_in_all_subarrays() {
	array := [37]int{1, 0, 1, 0, 2, 3, 0, 3, 2, 4, 0, 5, 1, 4, 5, 0, 1, 4, 0, 1, 3, 4, 0, 5, 4, 1, 7, 0, 5, 1, 1, 5, 4, 0, 5, 0, 0}
	window := 4
	i := 0
	j := window
	for j <= len(array) {
		find_zeros_count(array[i:j])
		i++
		j++
	}
}
func find_zeros_count(array []int) {
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			count++
		}
	}
	fmt.Printf("ARRAY:%v , count:%v", array, count)
	fmt.Println()
}

// Maximum difference between the first and last element in all windows of size k.
func Maximum_and_minimum_frist_and_last() {
	array := [37]int{1, 0, 1, 0, 2, 3, 0, 3, 2, 4, 0, 5, 1, 4, 5, 0, 1, 4, 0, 1, 3, 4, 0, 5, 4, 1, 7, 0, 5, 1, 1, 5, 4, 0, 5, 0, 0}
	window := 3
	i := 0
	j := window

	var Max int = 0
	var Mini int = 0
	for j <= len(array) {
		subarray := array[i:j]
		value := subarray[0] - subarray[len(subarray)-1]
		if value > Max {
			Max = value
		} else {
			if value < Mini {
				Mini = value
			}
		}
		i++
		j++
	}
	fmt.Println("Minimum:", Mini)
	fmt.Println("Maximum:", Max)
}

// Smallest subarray with sum ≥ target.
/*
1,2,3,4,5 - 5

1        [0:1]
1,2      [0:2]
1,2,3    [0:3]
1,2,3,4  [0:4]
1,2,3,4,5[0:5]

2      [1:2]
2,3    [1:2]
2,3,4  [1:3]
2,3,4,5[1:4]
*/
func Subarray() {
	array := [5]int{1, 2, 3, 4, 5}
	i := 0
	j := 0
	for i < len(array) {
		fmt.Println(array[i:j])
		i++
		j++
	}
}

package main

import (
	"fmt"
)

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

// Given a sorted array, remove all duplicates in place and return the new length.
/*
1 3 5 5 5
0 1 2 3 4
*/
func Remove_duplicates_in_place() {
	array := [5]int{5, 5, 1, 5, 3}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println("SORTED ARRAY:", array)
	i := 0
	j := 1
	for j < len(array) {
		if array[i] == array[j] {
			array[j] = 0
			j++
		} else {
			i++
			j++
		}
	}
	fmt.Println(array)
}

// Check if an array is a palindrome using two pointers.
/*
1,2,2,1
0,1,2,3
*/
func Palindrome() {
	array := []int{1, 2, 2, 1}
	i := 0
	j := len(array) - 1
	for i < j {
		if array[i] == array[j] {
			i++
			j--
		} else {
			fmt.Println("Not Palindrome!")
			return
		}
	}
	fmt.Println("Palindrome!")
}

// Reverse an array using the two-pointer technique.
func Reverse_array() {
	array := [5]int{1, 2, 3, 4, 5}
	i := 0
	j := len(array) - 1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
	fmt.Println(array)
}

// Find if a given pair of numbers in a sorted array sums to a target value.
/*
sorted array
Take some values that equal to the target value!
*/
func Pair_sum_equal_to_given_number() {
	slice := []int{1, 2, 3, 4, 5}
	sum_value := 4
	i := 0
	j := len(slice) - 1
	pairs := [][]int{}

	for i < j {
		sum := slice[i] + slice[j]

		if sum == sum_value {
			pairs = append(pairs, []int{slice[i], slice[j]})
			i++
			j-- // move both inward
		} else if sum < sum_value {
			i++ // need bigger sum
		} else {
			j-- // need smaller sum
		}
	}

	fmt.Println("Pairs with sum =", sum_value, ":", pairs)
}

/*
1,2,3,4,5
0,1,2,3,4

01
02
03
04
10
11
12
13
14
*/
func Using_the_Pairs() {
	slice := []int{1, 2, 3, 4, 5}
	sum_value := 4
	pairs := [][]int{}
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i]+slice[j] == sum_value {
				pair := []int{slice[i], slice[j]}
				pairs = append(pairs, pair)
			}

		}
	}
	fmt.Println(pairs)
}

// Move all zeros to the end of an array while maintaining the order of other elements.
/*
1,0,2,0,3,4,0,5 len := 8
0,1,2,3,4,5,6,7
*/
func Move_all_zeros_to_end() {
	slice := []int{1, 0, 2, 0, 3, 4, 0, 5}
	Move_all_zeros_to_end_using_for_loop(slice)

	fmt.Println("USING THE TWO POINTER TECHNIQUE")
	i := 0
	j := 1
	for j < len(slice)-1 {
		if slice[i] == 0 {
			slice[i], slice[j] = slice[j], slice[i]
		}
		i++
		j++
	}
	fmt.Println(slice)
}

// brute force
func Move_all_zeros_to_end_using_for_loop(slice []int) {
	fmt.Println("USING THE BRUTE FORCE ")

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == 0 {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	fmt.Println(slice)
}

// Separate even and odd numbers in an array using two pointers.
/*
1,2,3,4,5
0,1,2,3,4
*/
func Even_Odd() {
	slice := []int{1, 2, 3, 4, 5}
	even := []int{}
	odd := []int{}

	i := 0
	j := 1
	for j < len(slice) {
		fmt.Println(slice[i], slice[j])
		if slice[i]%2 == 0 {
			even = append(even, slice[i])
		} else {
			odd = append(odd, slice[i])
		}
		if slice[j]%2 == 0 {
			even = append(even, slice[j])
		} else {
			odd = append(odd, slice[j])
		}
		i = i + 2
		j = j + 2
	}
	fmt.Println("ODD:", odd)
	fmt.Println("EVEN:", even)
}

// Check if two strings are anagrams of each other using two pointers after sorting.

func Anagrams() {
	data1 := []int{1, 2, 3, 4, 5, 6}
	data2 := []int{5, 4, 1, 2, 3, 6}
	fmt.Println("ANAGRAMS USING FOR LOOP")
	Anagrams_Using_for_loop(data1, data2)
}

func Anagrams_Using_for_loop(array1 []int, array2 []int) {
	array_map := make(map[int]int)
	array_map2 := make(map[int]int)
	for key, _ := range array1 {
		array_map[key]++
	}
	for key, _ := range array2 {
		array_map2[key]++
	}
	fmt.Println(array_map)
	fmt.Println(array_map2)
	for keyone, valueone := range array_map {
		for keytwo, valuetwo := range array_map2 {
			if keyone == keytwo {
				if valueone == valuetwo {
					fmt.Println("ITS anagram!")
				}
			}
		}
	}
}

// Merge two sorted arrays into a single sorted array using two pointers.
func Merge_two_sorted_arrays_into_a_single_array() {
	fristarray := []int{1, 2, 3, 4, 5}
	secondarray := []int{2, 3, 4, 5, 6}
	new_array := []int{}
	i := 0
	j := 0
	for j < len(fristarray) || j < len(secondarray) {
		if fristarray[i] > secondarray[j] {
			new_slice := []int{secondarray[i], fristarray[i]}
			new_array = append(new_array, new_slice...)
		} else {
			new_slice := []int{fristarray[i], secondarray[i]}
			new_array = append(new_array, new_slice...)
		}
		i++
		j++
	}
	fmt.Println(new_array)
}
func Remove_duplicates_values() {
	data := "golang"
	rundata := []rune(data)
	for i := 0; i < len(rundata); i++ {
		for j := i + 1; j < len(rundata); j++ {
			if rundata[i] == rundata[j] {
				rundata[j] = 32
			}
		}
	}
	fmt.Println(string(rundata))
}

// Given a sorted array, find two numbers such that their sum equals a specific target.
func Specific_target_with_two_numbers() {
	array := [5]int{1, 2, 3, 4, 5}
	value := 4
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i]+array[j] == value {
				fmt.Println(array[i], array[j])
			}
		}
	}
}

// Find all unique pairs in an array whose sum equals a given number.

// Find the intersection of two sorted arrays using two pointers.

// Find the union of two sorted arrays using two pointers.

// Remove all duplicates from a sorted linked list using two pointers.

// Move all negative numbers to one side of the array using two pointers.

// Find the maximum product pair in an array using two pointers.

// Given a sorted array, find the pair whose sum is closest to zero.

// Given an array, check if it has any duplicate elements within k distance.

// Find the length of the longest substring without repeating characters using two pointers.
func main() {
	Specific_target_with_two_numbers()
}

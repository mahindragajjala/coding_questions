/*
Definition: A contiguous part of an array is called a subarray.
*/
package main

import "fmt"

func main() {
	Longest_subarray_with_sum()
}

// Basic
// Print all subarrays of an array
func Subarray(array [5]int) [][]int {
	var subarrays [][]int
	for i := 0; i < len(array); i++ {
		for k := i + 1; k <= len(array); k++ {
			subarrays = append(subarrays, array[i:k])
		}
	}
	return subarrays
}

// Count the number of subarrays in an array
func Count_number_of_subarray() {
	subarrays := Subarray([5]int{1, 2, 3, 4, 5})
	fmt.Println(len(subarrays))
}

// Fixed-size Subarrays
func Fixed_size_subarrays() [][]int {
	fixed_size := 3
	fixed_size_subarrays := [][]int{}
	array := []int{1, 2, 3, 4, 5}
	values := Subarray([5]int(array))
	for i := 0; i < len(values); i++ {
		if len(values[i]) == fixed_size {
			fmt.Println(values[i])
		}

	}
	return fixed_size_subarrays
}

// Maximum sum of a subarray of size k
func Maximum_sum_of_a_subarray_of_size_k() {
	array := []int{1, 2, 3, 4, 5}
	values := Subarray([5]int(array))
	for i := 0; i < len(values); i++ {
		values := Sum_of_subarray(values[i])
		if values == 0 {
			//--Don't Print the statement
		} else {
			fmt.Println(values)
		}

	}
}
func Sum(array []int) int {
	var Sum int
	for i := 0; i < len(array); i++ {
		Sum = Sum + array[i]
	}
	return Sum
}
func Sum_of_subarray(array []int) int {
	var k int = 3
	var Sum int
	if len(array) == k {
		fmt.Println(array)
		for i := 0; i < len(array); i++ {
			Sum = Sum + array[i]
		}
	}
	return Sum
}

// Minimum average of all subarrays of size k
/*
average := sum_of_subarrays / count of subarrays
*/
func Average_of_sub_arrays() int {
	var Average_of_sub_arrays int
	array := [5]int{1, 2, 3, 4, 5}
	var count int
	subarrays := Subarray(array)
	for i := 0; i < len(subarrays); i++ {
		count = count + Sum(subarrays[i])
	}
	Average_of_sub_arrays = count / len(subarrays)
	return Average_of_sub_arrays
}

// Maximum product of subarray of size k
func Product(array []int) int {
	var product int = 1
	for i := 0; i < len(array); i++ {
		product = product * array[i]
	}
	fmt.Println(product)
	return product
}
func Maximum_Product_of_subarray_of_size_k() int {
	var Maximum_Product int = 1
	array := [5]int{1, 2, 3, 4, 5}
	var k int = 3
	subarrays := Subarray(array)
	fmt.Println(subarrays)
	for i := 0; i < len(subarrays); i++ {
		if len(subarrays[i]) == k {
			fmt.Println(subarrays[i])
			product_value := Product(subarrays[i])
			if product_value > Maximum_Product {
				Maximum_Product = product_value
			}
		}
	}
	return Maximum_Product
}

// Variable-size Subarrays
func Variable_size_subarrays() {
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			subarray := array[i:j]
			fmt.Println(subarray)
		}
	}
}

// Smallest subarray with sum ≥ k
func Smallest_subarray_with_sum() {
	subarrays := Subarray([5]int{1, 2, 3, 4, 5})
	var sumofsubarray int
	var smallest_subarray []int
	for i := 0; i < len(subarrays); i++ {
		sumofarray := Sum(subarrays[i])
		if sumofarray < sumofsubarray {
			smallest_subarray = subarrays[i]
		}
	}
	fmt.Println(smallest_subarray)
}

// Longest subarray with sum ≤ k
func Longest_subarray_with_sum() {
	subarrays := Subarray([5]int{1, 2, 3, 4, 5})
	var sumofsubarray int
	var smallest_subarray []int
	for i := 0; i < len(subarrays); i++ {
		sumofarray := Sum(subarrays[i])
		fmt.Println(subarrays[i], sumofarray)
		if sumofarray > sumofsubarray {
			sumofsubarray = sumofarray
			smallest_subarray = subarrays[i]
		}
	}
	fmt.Println(smallest_subarray)
}

// Two Pointer + Sliding Window

// Subarrays with product less than k
// Longest subarray with all distinct elements
// Longest subarray with at most 2 distinct elements

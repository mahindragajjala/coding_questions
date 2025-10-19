package main

import "fmt"

//Maximum subarray sum (Kadane's Algorithm)
/*
Kadane’s Algorithm is a famous algorithm in computer science used to
find the maximum sum of a contiguous subarray within a one-dimensional
array of numbers (both positive and negative).
*/
func Maximum_subarray(slicedata []int) {
	var Maximum_sum int
	var slice_of_it []int
	for i := 0; i < len(slicedata); i++ {
		for j := i + 1; j < len(slicedata); j++ {
			var sum int
			for k := i; k < j; k++ {
				sum = sum + slicedata[k]
			}
			if sum > Maximum_sum {
				Maximum_sum = sum
				for f := i; f < j; f++ {
					slice_of_it = slice_of_it[:0]                        // reset slice
					slice_of_it = append(slice_of_it, slicedata[i:j]...) // store new subarray
				}
			}
		}
	}
	fmt.Println(Maximum_sum)
	fmt.Println(slice_of_it)
}

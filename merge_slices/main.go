package main

import "fmt"

//Merge two slices.
func main() {
	merge := Merge_slices([]int{1, 2, 3, 4, 5, 6}, []int{7, 8, 9, 0})
	fmt.Println(merge)
}

/*  */
func Merge_slices(one []int, two []int) []int {
	for i := 0; i < len(two); i++ {
		one = append(one, two[i])
	}
	return one
}

/*
OR
func Merge_slices(one []int, two []int) []int {
    return append(one, two...)
}
*/

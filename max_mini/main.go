package main

import "fmt"

/*
Find the maximum and minimum in a slice.
*/
func main() {
	slicedata := []int{1, 2, 3, 4, 5, 6}
	maximum := Max_Slice(slicedata)
	fmt.Println(maximum)
	minimum := Mini_slice(slicedata)
	fmt.Println(minimum)
}
func Max_Slice(slicedata []int) int {
	var max int = slicedata[0]
	for i := 0; i < len(slicedata); i++ {
		if slicedata[i] > max {
			max = slicedata[i]
		}
	}
	return max
}
func Mini_slice(slicedata []int) int {
	var mini int = slicedata[0]
	for i := 0; i < len(slicedata); i++ {
		if slicedata[i] < mini {
			mini = slicedata[i]
		}
	}
	return mini
}

package main

import "fmt"

/*
func main() {
	slicedata := []int{1, 3, 4, 8, 9, 2, 9}
	sort.Ints(slicedata)
	fmt.Println("Ascending:", slicedata)
}
*/

//Bubble/Selection Sort Logic
func main() {
	slicedata := []int{1, 3, 4, 8, 9, 2, 9}
	Ascending(slicedata)
}
func Ascending(slicedata []int) {
	for i := 0; i < len(slicedata); i++ {
		for j := i + 1; j < len(slicedata); j++ {
			if slicedata[j] < slicedata[i] {
				slicedata[i], slicedata[j] = slicedata[j], slicedata[i]
			}
		}
	}
	fmt.Println(slicedata)
}
func Decending(slicedata []int) {
	for i := 0; i < len(slicedata); i++ {
		for j := i + 1; j < len(slicedata); j++ {
			if slicedata[j] > slicedata[i] {
				slicedata[i], slicedata[j] = slicedata[j], slicedata[i]
			}
		}
	}
	fmt.Println(slicedata)
}

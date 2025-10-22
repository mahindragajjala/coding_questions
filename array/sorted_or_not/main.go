package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 44, 5, 6, 7}
	var max int = array[0]
	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		} else {
			fmt.Println("unsorted")
			return
		}
	}
	fmt.Println("sorted")
}

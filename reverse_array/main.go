package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
	fmt.Println(array)
}

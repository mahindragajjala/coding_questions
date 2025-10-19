package main

import "fmt"

func main() {
	//Find the average of elements
	array := [5]int{1, 2, 3, 4, 5}
	length := len(array)
	var sum int = 0
	var average int
	for i := 0; i < length; i++ {
		sum = sum + array[i]
	}

	average = sum / length
	fmt.Println(average)
}

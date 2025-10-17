package main

import "fmt"

/*
Implement recursion to count digits of a number.
*/
func main() {
	number := 23442342
	Count_Number(number)
}

/*
123

*/
func Count_Number(value int) int {
	var count int
	if value == 0 {
		fmt.Println(count)
		return count
	}
	value = value / 10
	count++
	return Count_Number(value)
}

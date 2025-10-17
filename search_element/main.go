package main

import "fmt"

func main() {
	data := []int{3, 2, 4, 8, 1, 3, 4, 1, 0, 9, 9, 8}
	value := 9
	indexes := Search_function(data, value)
	fmt.Println(indexes)
}
func Search_function(slice []int, value int) []int {
	searchdata_slice := []int{}
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			searchdata_slice = append(searchdata_slice, i+1)
		}
	}
	return searchdata_slice
}

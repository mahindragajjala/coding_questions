/*
Remove duplicates from a slice.
*/
package main

import "fmt"

func main() {
	Duplicates := []int{1, 2, 3, 1, 2, 3, 4, 1, 2, 4, 3, 1, 4, 1, 1, 3, 4, 5, 6, 7, 8, 9}
	nonduplicates := FindDuplicates(Duplicates)
	fmt.Println(nonduplicates)
}

func FindDuplicates(duplicates []int) (new []int) {
	mapdata := make(map[int]int)
	for i := 0; i < len(duplicates); i++ {
		mapdata[duplicates[i]]++
	}
	for key, value := range mapdata {
		if value <= 1 {
			new = append(new, key)
		}
	}
	return new
}

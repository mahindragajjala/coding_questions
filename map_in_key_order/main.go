package main

import (
	"fmt"
	"sort"
)

func main() {
	myMap := map[string]int{
		"banana": 3,
		"apple":  5,
		"cherry": 2,
		"mango":  7,
		"orange": 1,
	}
	strings_slice := []string{}
	for keys, _ := range myMap {
		strings_slice = append(strings_slice, keys)
	}
	sort.Strings(strings_slice)
	for i := 0; i < len(strings_slice); i++ {
		fmt.Println(strings_slice[i], myMap[strings_slice[i]])
	}
}

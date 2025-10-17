package main

import "fmt"

func main() {
	//Count the frequency of characters in a string using a map.

	stringdata := "golang"
	mapdata := make(map[string]int)
	for i := 0; i < len(stringdata); i++ {
		mapdata[string(stringdata[i])]++
	}
	fmt.Println(mapdata)
}

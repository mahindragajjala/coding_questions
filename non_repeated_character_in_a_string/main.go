package main

import "fmt"

func main() {
	stringdata := "string"

	// Step 1: create a map and a slice to preserve order
	mapdata := make(map[rune]int)
	order := []rune{}

	// Step 2: Count each character and store order
	for _, ch := range stringdata {
		mapdata[ch]++
		order = append(order, ch)
	}

	// Step 3: Find the first character with count == 1
	var firstNonRepeated rune
	for _, ch := range order {
		if mapdata[ch] == 1 {
			firstNonRepeated = ch
			break
		}
	}

	if firstNonRepeated != 0 {
		fmt.Printf("First non-repeated character: %c\n", firstNonRepeated)
	} else {
		fmt.Println("No non-repeated character found.")
	}
}

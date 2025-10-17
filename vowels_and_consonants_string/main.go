// Count the number of vowels and consonants in a string.
package main

import (
	"fmt"
	"strings"
)

func main() {
	stringdata := "Golang"
	lowercase := strings.ToLower(stringdata)
	var sliceOfdata []string
	var vowels, consonants int
	for i := 0; i < len(stringdata); i++ {
		sliceOfdata = append(sliceOfdata, string(lowercase[i]))
	}
	vowels, consonants = Vowels_or_Consonants(sliceOfdata)
	fmt.Printf("vowels:%v , consonants:%v", vowels, consonants)
}
func Vowels_or_Consonants(data []string) (vowels int, consonants int) {
	for i := 0; i < len(data); i++ {
		if data[i] == "a" || data[i] == "e" || data[i] == "i" || data[i] == "o" || data[i] == "u" {
			vowels++
		} else {
			consonants++
		}
	}

	return vowels, consonants
}

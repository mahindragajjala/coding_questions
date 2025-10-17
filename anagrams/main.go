package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to check if two strings are anagrams
func IsAnagram(str1, str2 string) bool {
	// Convert to lowercase for case-insensitive comparison
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// If lengths differ, cannot be anagrams
	if len(str1) != len(str2) {
		return false
	}

	// Convert strings to slices of runes and sort
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")
	sort.Strings(s1)
	sort.Strings(s2)

	// Compare sorted strings
	return strings.Join(s1, "") == strings.Join(s2, "")
}

func main() {
	word1 := "Listen"
	word2 := "Silent"

	if IsAnagram(word1, word2) {
		fmt.Println(word1, "and", word2, "are Anagrams ✅")
	} else {
		fmt.Println(word1, "and", word2, "are NOT Anagrams ❌")
	}
}

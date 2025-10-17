package main

import "fmt"

//Check if a string is a palindrome.
func main() {
	stringdata := "hello"
	stringslice := []rune(stringdata)
	for i := 0; i < len(stringdata)/2; i++ {
		stringslice[i], stringslice[len(stringslice)-1-i] = stringslice[len(stringslice)-1-i], stringslice[i]
	}
	reverse := string(stringslice)
	if stringdata == reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}

package main

import "fmt"

func main() {
	n := 432
	reverse := 0

	for n != 0 {
		digit := n % 10              // get last digit (2, then 3, then 4)
		reverse = reverse*10 + digit // add it to reverse number
		n = n / 10                   // remove last digit
	}

	fmt.Println("Reversed number:", reverse)
}

/*
1, 2, 3, 4, 5
0, 1  2, 3, 4,
*/

package main

import "fmt"

func main() {
	n := 10 // Number of Fibonacci numbers you want
	fib := Fibonacci(n)
	fmt.Println(fib)
}

func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	// Initialize first two numbers
	slice := []int{0, 1}

	// Generate remaining numbers
	for i := 2; i < n; i++ {
		next := slice[i-1] + slice[i-2]
		slice = append(slice, next)
	}

	return slice
}

package main

import (
	"math"
)

/*
Step 1:
Input a number n.

Step 2:
If n <= 1,
→ Not a prime (because prime numbers are greater than 1).

Step 3:
If n == 2,
→ Prime (since 2 is the smallest prime number).

Step 4:
If n is divisible by 2 (i.e., n % 2 == 0),
→ Not a prime (because all even numbers greater than 2 are not prime).

Step 5:
For all odd numbers i from 3 to √n (square root of n),
check:
→ If n % i == 0,
then n is not prime (because it has a factor other than 1 and itself).

Step 6:
If no such i divides n,
→ n is a prime number.
*/
func IsPrime(n int) bool {
	if n <= 1 {
		return false // 0 and 1 are not prime
	}
	if n == 2 {
		return true // 2 is prime
	}
	if n%2 == 0 {
		return false // even numbers >2 are not prime
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
i below 2 is not not prime
if 2 is prime
if num/2 == 0 is not prime
if number  to square root loop
*/

// ListPrimes returns all prime numbers up to n (inclusive)
func ListPrimes(n int) []int {
	primes := []int{}

	if n < 2 {
		return primes
	}

	for num := 2; num <= n; num++ {
		isPrime := true
		sqrtNum := int(math.Sqrt(float64(num)))

		for i := 2; i <= sqrtNum; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, num)
		}
	}

	return primes
}

package main

import "fmt"

func main() {
	n := 600851475143
	for i := 2; n >= i; i++ {
		if isPrime(i) {
			if n%i == 0 {
				n = n / i
				fmt.Println(i)
			}
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i < n; {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}

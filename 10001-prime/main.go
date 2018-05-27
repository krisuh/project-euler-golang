package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := parseInt(os.Args[1])
	pN(n)
}

func pN(n int) {
	var p int
	p = 0
	for i := 0; i < n; {
		if isPrime(p) {
			i++
			fmt.Println(i, p)
		}
		p++
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
	for i := 5; i*i <= n; {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

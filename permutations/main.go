package main

import (
	"fmt"
	"strings"
)

var count int

func main() {
	s := "0123456789"
	count = 0
	printPermutations("", s, len(s))
}

func printPermutations(s string, left string, length int) {
	if len(s) == length {
		count++
		if count == 1000000 {
			fmt.Println(s)
		}
	} else if count < 1000000 {
		for _, c := range left {
			printPermutations(s+string(c), strings.Replace(left, string(c), "", 1), length)
		}
	}
}

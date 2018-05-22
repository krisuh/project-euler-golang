package main

import (
	"fmt"
	"strconv"
)

func main() {
	max := 0
	var (
		maxL int
		maxR int
	)
	for r := 999; r > 100; r-- {
		for l := 998; l > 100; l-- {
			temp := l * r
			if isPalindrome(strconv.Itoa(temp)) {
				if max < temp {
					max = temp
					maxR = r
					maxL = l
				}
			}
		}
	}
	fmt.Println(max)
	fmt.Println(maxR)
	fmt.Println(maxL)
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

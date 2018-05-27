package main

import (
	"fmt"
	"strings"
)

func main() {
	printReplacedStrings("10055", "2")
}

func printReplacedStrings(orig string, with string) {
	for i := 0; i < len(orig); i++ {
		for j := i + 1; j < len(orig); j++ {
			temp := replaceAtIndex(orig, with, i)
			temp = replaceAtIndex(temp, with, j)
			fmt.Println(temp)
		}
	}
}

func replaceAtIndex(orig string, with string, index int) string {
	return strings.Join([]string{orig[:index], with, orig[index+1:]}, "")
}

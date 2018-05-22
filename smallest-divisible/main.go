package main

import "fmt"

func main() {
	found := false
	for i := 20; !found; i += 20 {
		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				break
			} else if i%j == 0 && j == 20 {
				found = true
				fmt.Println(i)
			}
		}
	}
}

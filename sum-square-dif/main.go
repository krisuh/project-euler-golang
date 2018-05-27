package main

import (
	"fmt"
	"math"
)

func main() {
	squareOfSums := math.Pow(arithmeticSeries(100, 1), 2.0)
	difference := squareOfSums - sumOfSquares(100)
	fmt.Println(difference)
}

func sumOfSquares(n float64) float64 {
	return (math.Pow(n, 3.0) / 3) + (math.Pow(n, 2.0) / 2.0) + (n / 6)
}

func arithmeticSeries(n float64, d float64) float64 {
	return n*1.0 + d*(((n-1.0)*n)/2.0)
}

/* 1, 4, 9, 16, 25, 36, 49, 64, 81, 100 */
/* 1, 3, 5, 7,  9,  11, 13, 15, 17, 19 */

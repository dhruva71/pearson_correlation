package main

import (
	"fmt"
	"pearson_correlation/correlation"
)

func main() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{5, 4, 3, 2, 1}
	fmt.Println(correlation.CalculateCorrelation(x, y))
}

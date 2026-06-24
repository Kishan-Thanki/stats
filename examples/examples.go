package main

import (
	"fmt"

	"github.com/kishan-thanki/stats"
)

func main() {
	intSlice := []int{1, 3, 3, 2, 3, 4, 5}
	avg, _ := stats.Average(intSlice)
	med, _ := stats.Median(intSlice)
	mode, _ := stats.Mode(intSlice)
	fmt.Printf("Stats on [1, 3, 3, 2, 3, 4, 5]: Avg: %.2f, Median: %.1f, Mode: %d\n", avg, med, mode)

	floatSlice := []float64{1.5, 3.5, 2.5}
	avgFloat, _ := stats.Average(floatSlice)
	fmt.Printf("Avg of [1.5, 3.5, 2.5]: %.2f\n", avgFloat)

	sumVal, _ := stats.SumSlice(intSlice)
	minVal, _ := stats.MinSlice(intSlice)
	maxVal, _ := stats.MaxSlice(intSlice)
	fmt.Printf("Slice utilities: Sum: %d, Min: %d, Max: %d\n", sumVal, minVal, maxVal)
}

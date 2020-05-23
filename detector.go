package main

import (
	"fmt"
	"math"
)

func findStdDev(arr []int) (float64, float64) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	mean := float64(sum) / float64(len(arr))

	sum2 := 0.0

	for i := 0; i < len(arr); i++ {
		sum2 += (float64(arr[i]) - mean) * (float64(arr[i]) - mean)
	}

	variance := sum2 / float64(len(arr))

	stdDev := math.Sqrt(variance)

	return mean, stdDev
}
func main() {
	arr := []int{1, 2, 3, 4, 5}

	mean, stdDev := findStdDev(arr)

	fmt.Println(mean, stdDev)

}

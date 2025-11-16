package main

import (
	"fmt"
	"slices"
)

func maxRain(arr []int) {
	n := len(arr)
	totalTrap := 0
	maxLeft := arr[0]
	maxRight := arr[n-1]

	maxLeftArray := make([]int, n)
	maxRightArray := make([]int, n)

	for i := 0; i < n; i++ {
		l := arr[i]
		r := arr[n-i-1]

		maxLeft = slices.Max([]int{l, maxLeft})
		maxRight = slices.Max([]int{r, maxRight})

		maxLeftArray[i] = maxLeft
		maxRightArray[n-i-1] = maxRight
	}
	for i := 0; i < n; i++ {
		totalTrap += slices.Min([]int{maxLeftArray[i], maxRightArray[i]}) - arr[i]
	}

	fmt.Println("Array:", arr)
	fmt.Println("Left :", maxLeftArray)
	fmt.Println("Right:", maxRightArray)
	fmt.Println("Total trapped water: ", totalTrap)
}

func main() {
	arr := [...]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	maxRain(arr[:])
}

package main

import (
	"fmt"
	"math"
	"slices"
)

func subArraySort(arr []int) {
	left := 0
	right := len(arr) - 1

	largest := math.MinInt16
	smallest := math.MaxInt16

	for i := 0; i < len(arr); i++ {
		if (i < len(arr)-1 && arr[i] > arr[i+1]) || (i > 0 && arr[i] < arr[i-1]) {
			largest = slices.Max([]int{largest, arr[i]})
			smallest = slices.Min([]int{smallest, arr[i]})
		}
	}

	fmt.Println("smallest, largest:", smallest, largest)

	for ; left < len(arr) && arr[left] <= smallest; left++ {
	}
	for ; right > 0 && arr[right] >= largest; right-- {
	}

	if left-right >= len(arr) {
		fmt.Println("Already sorted ", -1, -1)
		return
	}

	fmt.Println("Sub array sorted ", left, right)

}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 8, 6, 7, 9, 10, 11}

	subArraySort(arr[:])
}

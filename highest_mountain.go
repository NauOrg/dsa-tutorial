package main

import "fmt"

func highestMountain(arr []int) {
	maxMountains := 0

	if len(arr) <= 2 {
		maxMountains = len(arr)
	}

	for i := 1; i < len(arr)-1; i++ {

		if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
			c := 1

			for j := i; j > 0 && arr[j] > arr[j-1]; j, c = j-1, c+1 {
			}
			for ; i < len(arr)-1 && arr[i] > arr[i+1]; i, c = i+1, c+1 {
			}

			if c > maxMountains {
				maxMountains = c
			}
		}

	}

	fmt.Println("Longest Mountains: ", maxMountains)
}

func kadaneHeighestMountain(arr []int) {
	maxMountains := 0
	c := 0
	for i := 0; i < len(arr); i++ {
		c = c + 1
		if c > maxMountains {
			maxMountains = c
		}
		if i > 0 && i < len(arr)-1 && arr[i] < arr[i-1] && arr[i] < arr[i+1] {
			c = 1
		}
	}

	fmt.Println("Kadane Longest Mountains: ", maxMountains)
}

func main() {

	arr := [...]int{5, 6, 1, 2, 3, 4, 5, 4, 3, 2, 0, 1, 2, 3, -2, 4}
	highestMountain(arr[:])
	kadaneHeighestMountain(arr[:])
}

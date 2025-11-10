package main

import (
	"fmt"
	"math"
)

func closestSum(arr []int, target int) {

	minDiff := math.MaxInt32
	s := 0
	e := len(arr) - 1
	close1 := arr[0]
	close2 := arr[len(arr)-1]

	for s < e {
		sum := arr[s] + arr[e]
		diff := math.Abs(float64(sum) - float64(target))

		if diff < float64(minDiff) {
			minDiff = int(diff)
			close1 = arr[s]
			close2 = arr[e]
		}
		if sum < target {
			s++
		} else {
			e--
		}
	}

	fmt.Println(close1, close2, minDiff)
}

func main() {
	arr := [...]int{10, 22, 28, 29, 30, 40}
	target := 54
	closestSum(arr[:], target)
}

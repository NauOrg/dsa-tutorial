package main

import (
	"fmt"
)

func bruteForceMax(arr []int) {
	maxSum := 0

	for s := 0; s < len(arr); s++ {
		for e := s; e < len(arr); e++ {
			sum := 0
			for k := s; k <= e; k++ {
				sum += arr[k]
			}
			if maxSum < sum {
				maxSum = sum
			}
		}
	}

	fmt.Println("bruteForceMax", maxSum)
}

func prefixSum(arr []int) {
	maxSum := 0
	var prefixArr = make([]int, len(arr))

	prefixArr[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixArr[i] = prefixArr[i-1] + arr[i]
	}

	for s := 0; s < len(arr); s++ {
		for e := s; e < len(arr); e++ {
			sum := 0
			if s == 0 {
				sum = prefixArr[s]
			} else {

				sum = prefixArr[e] - prefixArr[s-1]
			}
			if maxSum < sum {
				maxSum = sum
			}
		}
	}

	fmt.Println("prefixSum", maxSum)
}

func kadaneMax(arr []int) {
	maxSum := 0
	sm := 0

	for i := 0; i < len(arr); i++ {
		sm = sm + arr[i]
		if sm < 0 {
			sm = 0
		}
		if sm > maxSum {
			maxSum = sm
		}
	}

	fmt.Println("kadeneMax", maxSum)
}

func main() {
	var arr = [...]int{1, -11, 3, -5, 7, 12, -9, 6, 4}
	bruteForceMax(arr[:])
	prefixSum(arr[:])
	kadeneMax(arr[:])
}

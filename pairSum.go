package main

import "fmt"

func main() {
	arr := [...]int{10, 5, 2, 3, -6, 9, 11}
	sum := 5

	Set := map[int]bool{}

	for _, v := range arr {

		if _, ok := Set[sum-v]; ok {
			fmt.Println(sum-v, v)
		} else {
			Set[v] = true
		}
	}

}

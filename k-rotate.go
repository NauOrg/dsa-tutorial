package main

import "fmt"

func kRotate(arr []int, k int) {
	newArray := make([]int, len(arr))
	j := 0
	n := len(arr)
	for i := n - k; i < n; i, j = i+1, j+1 {
		newArray[j] = arr[i]
	}
	for i := 0; i < n-k; i, j = i+1, j+1 {
		newArray[j] = arr[i]
	}

	fmt.Println("ROTATED ARRAY", newArray)
}
func kRotateMod(arr []int, k int) {
	newArray := make([]int, len(arr))

	n := len(arr)
	for i := 0; i < n; i = i + 1 {
		newArray[(i+k)%n] = arr[i]
	}

	fmt.Println("MOD ROTATED ARRAY", newArray)
}

func reverse(arr []int) {
	n := len(arr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func kRotateReverse(arr []int, k int) {

	n := len(arr)
	reverse(arr[n-k:])
	reverse(arr[:n-k])
	reverse(arr)

	fmt.Println("Reverse ROTATED ARRAY", arr)
}

func main() {

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println("Original: ", arr)
	kRotate(arr[:], k)
	kRotateMod(arr[:], k)
	kRotateReverse(arr[:], k)

}

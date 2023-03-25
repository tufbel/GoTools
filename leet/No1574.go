// Package leet
// Title       : No1574.go
// Author      : Tuffy  2023/3/25 10:58
// Description :
package leet

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	j := n - 1
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if j == 0 {
		return 0
	}
	res := j
	for i := 0; i < n; i++ {
		for j < n && arr[j] < arr[i] {
			j++
		}
		res = min(res, j-i-1)
		if i+1 < n && arr[i] > arr[i+1] {
			break
		}
	}
	return res
}

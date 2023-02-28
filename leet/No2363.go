// Package leet
// Title       : No2363.go
// Author      : Tuffy  2023/2/28 13:38
// Description :
package leet

import "sort"

func mergeSimilarItems(item1, item2 [][]int) [][]int {
	mp := map[int]int{}
	for _, a := range item1 {
		mp[a[0]] += a[1]
	}
	for _, a := range item2 {
		mp[a[0]] += a[1]
	}
	var ans [][]int
	for a, b := range mp {
		ans = append(ans, []int{a, b})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

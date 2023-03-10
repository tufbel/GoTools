// Package leet
// Title       : No1590.go
// Author      : Tuffy  2023/3/10 15:28
// Description :
package leet

func minSubarray(nums []int, p int) int {
	sum := 0
	mp := map[int]int{0: -1}
	for _, v := range nums {
		sum += v
	}
	rem := sum % p
	if rem == 0 {
		return 0
	}
	minCount := len(nums)
	sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		tempRem := sum % p
		k := (tempRem - rem + p) % p
		if _, ok := mp[k]; ok {
			minCount = min(minCount, i-mp[k])
		}
		mp[tempRem] = i
	}

	if minCount >= len(nums) {
		return -1
	}

	return minCount
}

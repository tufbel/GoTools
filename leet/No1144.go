// Package leet
// Title       : No1144.go
// Author      : Tuffy  2023/2/27 16:50
// Description :
package leet

func movesToMakeZigzag(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	odd_operand := 0
	even_operand := 0
	nums = append([]int{nums[1]}, nums...)
	nums = append(nums, nums[len(nums)-2])
	for i := 1; i < len(nums)-1; i++ {
		targetNum := min(nums[i-1], nums[i+1]) - 1
		if nums[i] > targetNum {
			switch i % 2 {
			case 0:
				even_operand += nums[i] - targetNum
			case 1:
				odd_operand += nums[i] - targetNum
			}
		}
	}
	return min(odd_operand, even_operand)
}

//func init() {
//	println(movesToMakeZigzag([]int{9, 6, 1, 6, 2}))
//}

package leet

// 从后往前遍历， 单调栈寻找第一个大于当前值的数

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := map[int]int{}
	decrStack := []int{}
	for i := len(nums2) - 1; i > 0; i-- {
		num := nums2[i]
		for len(decrStack) > 0 && num >= decrStack[len(decrStack)-1] {
			decrStack = decrStack[:len(decrStack)-1]
		}
		if len(decrStack) > 0 {
			res[num] = decrStack[len(decrStack)-1]
		} else {
			res[num] = -1
		}
		decrStack = append(decrStack, num)
	}
	ans := make([]int, len(nums1))
	for i, val := range nums1 {
		ans[i] = res[val]
	}
	return ans
}

package leet

// 计算出所有天数的劳累程度，利用递减单调栈寻找增区间

func longestWPI(hours []int) (ans int) {
	work_days := len(hours)
	index_sum := make([]int, work_days+1) // 前缀和
	decrStack := []int{0}
	for i, hour := range hours {
		if hour > 8 {
			index_sum[i+1] = index_sum[i] + 1
		} else {
			index_sum[i+1] = index_sum[i] - 1
		}
		if index_sum[i+1] < index_sum[decrStack[len(decrStack)-1]] {
			decrStack = append(decrStack, i+1)
		}
	}
	for i := work_days; i > 0; i-- {
		for len(decrStack) > 0 && index_sum[i] > index_sum[decrStack[len(decrStack)-1]] {
			ans = max(ans, i-decrStack[len(decrStack)-1])
			decrStack = decrStack[:len(decrStack)-1]
		}
	}
	return
}

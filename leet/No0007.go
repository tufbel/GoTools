package leet

import "math"

// 每次对10取余，弹出尾部，然后旧的值*10 + 当前尾部数
// go 对 附属取摸取得是 此附属相反数取余后的相反数， 如果想得到真正的取模应该是  value % n + n

func reverse(x int) (return_num int) {
	for x != 0 {
		if return_num < math.MinInt32/10 || return_num > math.MaxInt32/10 {
			return 0
		}
		tmp := x % 10
		x /= 10
		return_num = return_num*10 + tmp
	}
	return
}

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 示例演示主动抛出一个panic, 以及panic后进行recover恢复，并指定返回值的实例
func parseInt(str string) (ret int, retError error) {
	defer func() {
		if r := recover(); r != nil {
			ret = -1
			//retError = errors.New("this is a specified ")
		}
	}()
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing '%s': %v", str, err)
	}
	panic("raise exception")
	fmt.Println("resume exec")
	return int(i), errors.New("this is error")
}

func main() {
	value, err := parseInt("20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

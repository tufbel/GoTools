package reflect_eg

import (
	"fmt"
	"reflect"
)

func TypeAsserRun() {
	var x interface{} = "hello"

	s, ok := x.(string) // 类型断言并转换，ok为true时，s为转换后的值，否则为类型的零值
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("x is not a string")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
			fmt.Println("reflect x type:", reflect.TypeOf(x))
		}
	}()

	val := x.(string)
	println("val to string:", val)

	v := x.(int) // 类型断言失败，会panic
	println(v)
}

func SwitchType() {
	var x interface{} = []int{1, 2, 3}
	switch v := x.(type) { // x.(type)只能在switch语句中使用
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("v type is:", reflect.TypeOf(x))
	}
}

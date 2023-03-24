// Package struct_and_class
// Title       : equal_struct.go
// Author      : Tuffy  2023/3/11 9:57
// Description :
package struct_and_class

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

type Person2 struct {
	Name    string
	Age     int
	Address *Address
}

type Address struct {
	City   string
	Street string
}

func equal_struct_init() {
	p1 := Person{Name: "Alice", Age: 30, Gender: "Female"}
	p2 := Person{Name: "Alice", Age: 30, Gender: "Female"}
	p3 := Person{Name: "Bob", Age: 25, Gender: "Male"}

	fmt.Printf("p1 == p2 :%v\n", p1 == p2) // 输出 true
	fmt.Printf("p1 == p3 :%v\n", p1 == p3) // 输出 false

	tp1 := Person2{
		Name:    "Alice",
		Age:     30,
		Address: &Address{City: "Beijing", Street: "Main St."},
	}

	tp2 := Person2{
		Name:    "Alice",
		Age:     30,
		Address: &Address{City: "Beijing", Street: "Main St."},
	}

	fmt.Printf("tp1 == tp2 :%v", reflect.DeepEqual(tp1, tp2)) // 输出 true
}

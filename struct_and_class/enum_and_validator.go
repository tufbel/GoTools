// Package struct_and_class
// Title       : enum_and_validator.go
// Author      : Tuffy  2023/3/24 16:53
// Description :
package struct_and_class

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	Apple  string = "apple"
	Orange string = "orange"
	Banana string = "banana"
)

type MyStruct struct {
	//Fruit string `validate:"required"`
	Fruit string `validate:"required,oneof=apple orange banana"`
}

func enum_and_validator_init() {
	validate := validator.New()
	myStruct := MyStruct{Fruit: "cherry"}

	err := validate.Struct(myStruct)
	if err != nil {
		fmt.Println(err) // 输出：Key: 'MyStruct.Fruit' Error:Field validation for 'Fruit' failed on the 'oneof' tag
	}
	fmt.Printf("%+v\n", myStruct)
}

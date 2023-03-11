package struct_and_class

import "fmt"

type MethodStruct struct {
	name   string
	age    int
	myType string
}

func (*MethodStruct) newMethodStruct(
	name string,
	opts ...func(*MethodStruct),
) *MethodStruct {
	ncla := &MethodStruct{
		name:   name,
		age:    18,
		myType: "T_Type",
	}
	for _, opt := range opts {
		opt(ncla)
	}
	return ncla
}

func (*MethodStruct) ageOpt(age int) func(*MethodStruct) {
	return func(methodStruct *MethodStruct) {
		methodStruct.age = age
	}
}

func (*MethodStruct) myTypeOpt(myType string) func(*MethodStruct) {
	return func(methodStruct *MethodStruct) {
		methodStruct.myType = myType
	}
}

func (*MethodStruct) myStaticMethod(name string) string {
	return fmt.Sprintf("MethodSturct static method format %s", name)
}

func (receiver *MethodStruct) setMyType(myType string) {
	receiver.myType = myType
}

func StaticMethodMain() {
	var m *MethodStruct
	//m.setMyType("Type")
	fmt.Printf("Before using set method %+v\n", m)
	m = m.newMethodStruct("tuffy", m.ageOpt(20))
	fmt.Printf("New struct is %+v\n", m)
	fmt.Printf("Static method return: %s\n", m.myStaticMethod("Jerry"))
	fmt.Printf("After using static methods %+v\n", m)
	m.setMyType("G_Type")
	fmt.Printf("After using set method %+v\n", m)
	(*m).setMyType("**_Type")
	fmt.Printf("After both  using set method %+v\n", m)
}

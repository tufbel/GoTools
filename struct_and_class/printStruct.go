package struct_and_class

import "fmt"

type TuffyStruct struct {
	name string
	age  int
}

func (s TuffyStruct) String() string {
	return "<TuffyStruct>"
}

func main() {
	ts := TuffyStruct{"Tuffy", 18}
	fmt.Printf("My struct is %s\n", ts)
	fmt.Printf("My struct is %v\n", ts)
	fmt.Printf("My struct is %+v\n", ts)
	fmt.Printf("My struct is %#v\n", ts)
	fmt.Printf("My struct is %T\n", ts)
	fmt.Printf("My struct is %%\n")
	fmt.Println("hello", "world")
}

// Package learning
// Title       : reference_type.go
// Author      : Tuffy  2023/3/1 15:50
// Description :
package learning

import "fmt"

func referenceTypeRun() {
	myPoninter := map[string]string{"name": "tuffy"}
	myString := "Tuffy"
	fmt.Printf("myPoninter: %+v\n", myPoninter)
	fmt.Printf("myPoninter: %p\n", myPoninter)
	fmt.Printf("myPoninter: %p\n", &myPoninter)
	fmt.Printf("myPoninter: %T\n", myPoninter)
	fmt.Printf("myPoninter: %p\n", myString)
	fmt.Printf("myPoninter: %p\n", &myString)
}

//func init() {
//	referenceTypeRun()
//}

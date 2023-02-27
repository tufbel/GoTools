package main

import (
	"fmt"
	"reflect"
)

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	bubbleArray := []int{3, 5, 2, 1, 2, 3, 8}

	myArray := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(myArray)) // 5
	fmt.Printf("Type of x is %s\n", reflect.TypeOf(myArray))

	BubbleSort(bubbleArray)
}

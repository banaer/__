package main

import (
	"fmt"

)

// 固定长度的是值传递，不固定长度的是引用传递
func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("change %v %p\n", a, &a)
	change(a)
	fmt.Printf("change %v %p\n", a, &a)
	fmt.Printf("%d\n", cap(a))

	// fmt.Printf("changeWithRange %v %p\n", a, &a)
	// changeWithRange(a)
	// fmt.Printf("changeWithRange %v %p\n", a, &a)
}

func change(arr []int) {
	fmt.Printf("add = %p", &arr)
	for k, _ := range arr {
		arr[k] = 0
	}
}

func changeWithRange(arr [8]int) {
	fmt.Printf("add = %p", &arr)
	for k, _ := range arr {
		arr[k] = 0
	}
}

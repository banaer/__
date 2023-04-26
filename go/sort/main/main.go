package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[:4]
	fmt.Println(a)
	fmt.Println(b[1:])
}

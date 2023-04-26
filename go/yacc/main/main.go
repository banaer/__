package main

import (
	"bytes"
	"compute"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(([]byte)("12345671234567"))
	c := compute.Lex{Input: b}
	fmt.Println("------",c)
	compute.Parse(&c)
	fmt.Println("------",c)
}

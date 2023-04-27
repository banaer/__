package main

import (
	"bytes"
	"compute"
)

func main() {
	compute.Parse(&compute.Lex{Input: bytes.NewBuffer(([]byte)("12345  +    671234567   "))})
	compute.Parse(&compute.Lex{Input: bytes.NewBuffer(([]byte)("12345  -    671234567"))})
	compute.Parse(&compute.Lex{Input: bytes.NewBuffer(([]byte)("12345  *    671234567"))})
	compute.Parse(&compute.Lex{Input: bytes.NewBuffer(([]byte)("12345  /    671234567"))})
	compute.Parse(&compute.Lex{Input: bytes.NewBuffer(([]byte)("12345  %    671234567"))})
}

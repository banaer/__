package compute

import (
	"bytes"
	"fmt"
	"unicode"
)

type Compute struct {
	output string
}

type Lex struct {
	Input   *bytes.Buffer
	compute Compute
}

func (c *Lex) Lex(lval *yySymType) int {
	b, err := c.Input.ReadByte()
	if err == nil {
		lval.ele = string(b)
		if unicode.IsNumber(rune(b)) {
			return NUM
		}
		if unicode.IsSpace(rune(b)) {
			return SPACE
		}
		if bytes.IndexByte([]byte("+-*/%"), b) > -1 {
			return OPT
		}
		return 0
	}
	return 0
}

func (c *Lex) Error(s string) {
	fmt.Println("error", s)
}

func Parse(c *Lex) {
	yyParse(c)
}

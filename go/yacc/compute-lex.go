package compute

import (
	"bytes"
	"fmt"
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
		fmt.Println("-----", lval.ele, lval.str)
		return NUM
	}
	fmt.Println("errrr", err, b)
	return 0
}

func (c *Lex) Error(s string) {
	fmt.Println("error", s)
}

func Parse(c *Lex) {
	yyParse(c)
}

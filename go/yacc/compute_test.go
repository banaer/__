package compute

import "testing"

func TestCompute(
	t *testing.T) {
	c := Compute{}
	yyParse(c)
	t.Log(c.result)
}

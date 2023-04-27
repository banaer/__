%{
package compute
import "fmt"
import "strconv"
func setResult(l yyLexer, v Compute){
  fmt.Println("output",v.output)
  l.(*Lex).compute = v
}
%}

%union {
  compute Compute
  str string
  ele string
}

%token NUM OPT SPACE
%type <str> expr
%type <ele> NUM n expr1 op OPT
%start main

%%

main: expr
  {
    setResult(yylex, Compute{output:$1})
  }
expr:
  expr1 op expr1
  {
    ret := ""
    a,_ := strconv.Atoi($1)
    b,_ := strconv.Atoi($3)
    switch $2 {
      case "+":
        ret = fmt.Sprint(a + b )
      case "-":
        ret = fmt.Sprint(a - b )
      case "*":
        ret = fmt.Sprint(a * b )
      case "/":
        ret = fmt.Sprint(a / b )
      case "%":
        ret = fmt.Sprint(a % b )
    }
    $$ = $1 + $2 + $3 + "=" + ret
  }
op:
  OPT
  {
    $$ = $1
  }
| op SPACE
  {
    $$ = $1
  }
expr1:
  expr1 SPACE
  {
    $$ = $1
  }
|  n
  {
    $$ = $1
  }
n:
  NUM
  {
    $$ = $1
  }
| n NUM
  {
    $$ = $1 + $2
  }
%{
package compute
import "fmt"
func setResult(l yyLexer, v Compute){
  fmt.Println(v)
  l.(*Lex).compute = v
  // l.(*Compute).output = v.output
  // l.(*Compute).result = result
}
%}

%union {
  compute Compute
  str string
  ele string
}

%token NUM
%type <str> expr
%type <ele> NUM n
%start main

%%

main: expr
  {
    setResult(yylex, Compute{output:$1})
  }
expr:
  n
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
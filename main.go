package main

import (
  "os"

  "RETIA/unit"
  "RETIA/show"
)

type Result struct {
  data string
}

func (result *Result) Query(tuple *unit.Tuple) {
  if tuple != nil {
    show.Tuple(tuple)
  }
}


func main() {
  yyErrorVerbose = true

  result := new(Result)

  for {
    yyParse(NewLexerWithInit(os.Stdin, func(y *Lexer) { y.p = result }))
  }
}


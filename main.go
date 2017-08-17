package main

import (
  "os"

  "RETIA/unit"
  "RETIA/show"
)

type Result struct {
  data string
}

func (result *Result) Query(tuple *unit.Tuple, relation *unit.Relation) {
  if tuple != nil {
    show.Tuple(tuple)
  }

  if relation != nil {
    show.Relation(relation)
  }
}


func main() {
  yyErrorVerbose = true

  result := new(Result)

  for {
    yyParse(NewLexerWithInit(os.Stdin, func(y *Lexer) { y.p = result }))
  }
}


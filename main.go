package main

import (
  "fmt"
  "os"
  // "strconv"

  "RETIA/unit"
)

type Result struct {
  data string
}

func (result *Result) Query(tuple *unit.Tuple) {
  if tuple != nil {
    fmt.Printf("TUPLE ")

    if len(tuple.Vname) != 0 {
      fmt.Printf("(%s) ", tuple.Vname)
    }

    fmt.Printf("[%s] ", tuple.Tname)

    fmt.Printf("{ \n")

    if tuple.Components != nil {
      for _, component := range tuple.Components {
        fmt.Printf("        (%s %s %s) \n", component.Aname, component.Atype, component.Cvalue)
      }
    }

    fmt.Printf("      } \n")
  }
}


/* func atoi(s string) int64 {
  n, e := strconv.Atoi(s)
  if e != nil {
    panic(e)
  }
  return int64(n)
} */


func main() {
  yyErrorVerbose = true

  result := new(Result)

  for {
    yyParse(NewLexerWithInit(os.Stdin, func(y *Lexer) { y.p = result }))
  }
}


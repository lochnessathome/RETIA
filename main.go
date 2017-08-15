package main

import (
  "fmt"
  "os"
  "strconv"
)


type Result struct {
  data string
}


type Tuple struct {
  attributes []*Attribute
  components []*Component
}


type Attribute struct {
  aname string
  atype string
}

type Component struct {
  cname string
  ctype string
  cvalue string
}


func NewAttribute(aname, atype string) *Attribute {
  return &Attribute{aname, atype}
}

func NewComponent(cname, ctype, cvalue string) *Component {
  return &Component{cname, ctype, cvalue}
}

func NewTupleHeading(attributes []*Attribute) *Tuple {
  return &Tuple{attributes, nil}
}

func NewTupleBody(components []*Component) *Tuple {
  return &Tuple{nil, components}
}


func (result *Result) Query(tuple *Tuple) {
  fmt.Printf("TUPLE { \n")

  if tuple.attributes != nil {
    for _, attribute := range tuple.attributes {
      fmt.Printf("        (%s %s) \n", attribute.aname, attribute.atype)
    }
  }

  if tuple.components != nil {
    for _, component := range tuple.components {
      fmt.Printf("        (%s %s %s) \n", component.cname, component.ctype, component.cvalue)
    }
  }

  fmt.Printf("      } \n")
}


func atoi(s string) int64 {
  n, e := strconv.Atoi(s)
  if e != nil {
    panic(e)
  }
  return int64(n)
}


func main() {
  result := new(Result)

  yyParse(NewLexerWithInit(os.Stdin, func(y *Lexer) { y.p = result }))
}


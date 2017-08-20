package unit

import (
  "strings"
)


type Component struct {
  Aname string
  Atype string
  Cvalue string
}

type Tuple struct {
  Vname string
  Tname string
  Components []*Component
}

type Relation struct {
  Vname string
  Tname string
  Tuples []*Tuple
}


type ReductionStatement struct {
  Vname string
  Relation *Relation
  CompareExpression *CompareExpression
}

type CompareExpression struct {
  Laname string
  Raname string
  Rcvalue string
  Rctype string
  Operator string
}


func FormatLetter(letter string) (string) {
  return strings.ToLower(letter)  
}

func FormatTypeStr(aname, atype string) (string) {
  return ("[" + aname + "=" + atype + "]")
}

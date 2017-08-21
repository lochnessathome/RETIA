package unit

import (
  "strings"
  "crypto/md5"
  "encoding/hex"
)


type Component struct {
  Aname string
  Atype string
  Cvalue string
}

type Tuple struct {
  Vname string
  Tname string
  Hash string
  Components []*Component
}

type Relation struct {
  Vname string
  Tname string
  Tuples []*Tuple
}


type ReductionStatement struct {
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

type UnionStatement struct {
  Lrelation *Relation
  Rrelation *Relation
}

type IntersectionStatement struct {
  Lrelation *Relation
  Rrelation *Relation
}

type MinusStatement struct {
  Lrelation *Relation
  Rrelation *Relation
}

type TimesStatement struct {
  Lrelation *Relation
  Rrelation *Relation
}

type JoinStatement struct {
  Lrelation *Relation
  Rrelation *Relation
}


func FormatLetter(letter string) (string) {
  return strings.ToLower(letter)  
}

func FormatTypeStr(aname, atype string) (string) {
  return ("[" + aname + "=" + atype + "]")
}

func FormatHash(text string) (string) {
  hash := md5.Sum([]byte(text))
  return hex.EncodeToString(hash[:])
}


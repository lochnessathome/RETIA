package messages

import (
  "fmt"

  "RETIA/unit"
)


func Tuple(tuple *unit.Tuple) {
  showTuple(tuple, "")
}

func Relation(relation *unit.Relation) {
  fmt.Printf("RELATION ")

  if len(relation.Vname) != 0 {
    fmt.Printf("(%s) ", relation.Vname)
  }

  // fmt.Printf("[%s] ", relation.Tname)

  fmt.Printf("{ \n")

  if relation.Tuples != nil {
    for _, tuple := range relation.Tuples {
      showTuple(tuple, "         ")
    }
  }

  fmt.Printf("         } \n")
}

func VnameBusy(vname string) {
  fmt.Printf("Variable %s already exists.\n", vname)
}

func VnameMissing(vname string) {
  fmt.Printf("Variable %s not found.\n", vname)
}

func TypesMismatch() {
  fmt.Printf("Given relations have different types.\n")
}


func showTuple(tuple *unit.Tuple, prefix string) {
  fmt.Printf("%sTUPLE ", prefix)

  if len(tuple.Vname) != 0 {
    fmt.Printf("(%s) ", tuple.Vname)
  }

  // fmt.Printf("[%s] ", tuple.Tname)

  fmt.Printf("{ \n")

  if tuple.Components != nil {
    for _, component := range tuple.Components {
      fmt.Printf("%s        (%s %s %s) \n", prefix, component.Aname, component.Atype, component.Cvalue)
    }
  }

  fmt.Printf("%s      } \n", prefix)
}


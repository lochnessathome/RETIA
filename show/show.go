package show

import (
  "fmt"

  "RETIA/unit"
)


func Tuple(tuple *unit.Tuple) {
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


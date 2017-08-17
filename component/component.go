package component

import (
  "fmt"

  "RETIA/unit"
)


func Create(aname, atype, cvalue, ctype string) *unit.Component {
  atype = unit.FormatLetter(atype)

  if typeMatches(atype, ctype) {
    component := new(unit.Component)

    component.Aname = unit.FormatLetter(aname)
    component.Atype = atype
    component.Cvalue = cvalue

    return component
  } else {
    fmt.Printf("Given component %s has type %s. Expected type is %s. \n", cvalue, ctype, atype)

    return nil
  }
}

func typeMatches(atype, ctype string) bool {
  return (atype == ctype)
}


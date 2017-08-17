package component

import (
  "strings"
  "fmt"

  "RETIA/unit"
)


func Create(aname, atype, cvalue, ctype string) *unit.Component {
  if typeMatches(atype, ctype) {
    component := new(unit.Component)

    component.Aname = aname
    component.Atype = atype
    component.Cvalue = cvalue

    return component
  } else {
    fmt.Printf("%s should be type of %s.\n", cvalue, strings.ToLower(atype))

    return nil
  }
}

func typeMatches(atype, ctype string) bool {
  return (strings.ToLower(atype) == strings.ToLower(ctype))
}


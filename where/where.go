package where

import (
  "fmt"

  "RETIA/unit"
)


func Create(relation *unit.Relation, compare *unit.Compare) *unit.Where {
  if attributesValid(relation, compare) {
    where := new(unit.Where)

    where.Relation = relation
    where.Compare = compare

    return where
  } else {
    return nil
  }
}


func attributesValid(relation *unit.Relation, compare *unit.Compare) bool {
  loptype := ""
  roptype := ""

  _, loptype = findAttrByName(relation, compare.Laname)

  if len(loptype) == 0 {
    fmt.Printf("Given attribute %s not found in relation %s. \n", compare.Laname, relation.Vname)

    return false
  }

  if len(compare.Raname) != 0 {
    _, roptype = findAttrByName(relation, compare.Raname)

    if len(roptype) == 0 {
      fmt.Printf("Given attribute %s not found in relation %s. \n", compare.Raname, relation.Vname)

      return false
    }
  } else {
    roptype = compare.Rctype
  }

  if loptype != roptype {
    fmt.Printf("Can't compare values of different types: %s and %s. \n", loptype, roptype)

    return false
  }

  return true
}

func findAttrByName(relation *unit.Relation, aname string) (string, string) {
  for _, component := range relation.Tuples[0].Components {
    if component.Aname == aname {
      return component.Aname, component.Atype
    }
  }

  return "", ""
}


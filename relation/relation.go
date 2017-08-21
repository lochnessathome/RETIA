package relation

import (
  "fmt"

  "RETIA/unit"
)


func Create(tuples []*unit.Tuple) *unit.Relation {
  if len(tuples) != 0 && tuplesValid(tuples, tuples[0].Tname) {
    relation := new(unit.Relation)

    relation.Tname = tuples[0].Tname

    relation.Tuples = tuples
    relation.Vname = "" /* unit.FormatLetter() */

    return relation
  } else {
    return nil
  }
}

func tuplesValid(tuples []*unit.Tuple, base_tname string) bool {
  for _, tuple := range tuples {
    if tuple == nil {
      return false
    }

    if tuple.Tname != base_tname {
      fmt.Printf("Given tuple has type %s. Expected type is %s. \n", tuple.Tname, base_tname)

      return false
    }
  }

  return true
}


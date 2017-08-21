package minus

import (
  "RETIA/unit"
)


func Create(lrelation, rrelation *unit.Relation) *unit.MinusStatement {
  if lrelation != nil && rrelation != nil && relationsTypeMatches(lrelation, rrelation) {
    statement := new(unit.MinusStatement)

    statement.Lrelation = lrelation
    statement.Rrelation = rrelation

    return statement
  } else {
    return nil
  }
}


func Eval(statement *unit.MinusStatement) *unit.Relation {
  relation := new(unit.Relation)

  relation.Tname = statement.Lrelation.Tname

  for _, l_tuple := range statement.Lrelation.Tuples {
    present := false

    for _, r_tuple := range statement.Rrelation.Tuples {
      if l_tuple.Hash == r_tuple.Hash {
        present = true
        break
      }
    }

    if !present {
      relation.Tuples = append(relation.Tuples, l_tuple)
    }
  }

  return relation
}


func relationsTypeMatches(lrelation, rrelation *unit.Relation) bool {
  return (lrelation.Tname == rrelation.Tname)
}


package intersection

import (
  "RETIA/unit"
)


func Create(lrelation, rrelation *unit.Relation) *unit.IntersectionStatement {
  if lrelation != nil && rrelation != nil && relationsTypeMatches(lrelation, rrelation) {
    statement := new(unit.IntersectionStatement)

    statement.Lrelation = lrelation
    statement.Rrelation = rrelation

    return statement
  } else {
    return nil
  }
}


func Eval(statement *unit.IntersectionStatement) *unit.Relation {
  relation := new(unit.Relation)

  relation.Tname = statement.Lrelation.Tname

  for _, r_tuple := range statement.Rrelation.Tuples {
    for _, l_tuple := range statement.Lrelation.Tuples {
      if r_tuple.Hash == l_tuple.Hash {
        relation.Tuples = append(relation.Tuples, r_tuple)
        break
      }
    }
  }

  return relation
}


func relationsTypeMatches(lrelation, rrelation *unit.Relation) bool {
  return (lrelation.Tname == rrelation.Tname)
}


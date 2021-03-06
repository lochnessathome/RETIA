package intersection

import (
  "RETIA/unit"
  "RETIA/messages"
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
  if statement != nil {
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
  } else {
    return nil
  }
}


func relationsTypeMatches(lrelation, rrelation *unit.Relation) bool {
  if lrelation.Tname == rrelation.Tname {
    return true
  } else {
    messages.TypesMismatch()
    return true
  }
}


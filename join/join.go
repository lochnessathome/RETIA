package join

import (
  "RETIA/unit"
)


func Create(lrelation, rrelation *unit.Relation, vname string) *unit.JoinStatement {
  if lrelation != nil && rrelation != nil && relationsTypeMatches(lrelation, rrelation) {
    statement := new(unit.JoinStatement)

    statement.Lrelation = lrelation
    statement.Rrelation = rrelation

    statement.Vname = unit.FormatLetter(vname)

    return statement
  } else {
    return nil
  }
}


func Eval(statement *unit.JoinStatement) *unit.Relation {
  relation := new(unit.Relation)

  relation.Tname = statement.Lrelation.Tname
  relation.Vname = statement.Vname

  // TODO: how to copy an array?

  for _, l_tuple := range statement.Lrelation.Tuples {
    relation.Tuples = append(relation.Tuples, l_tuple)
  }

  for _, r_tuple := range statement.Rrelation.Tuples {
    matches := false

    for _, l_tuple := range statement.Lrelation.Tuples {
      if r_tuple.Hash == l_tuple.Hash {
        matches = true
        break
      }
    }

    if !matches {
      relation.Tuples = append(relation.Tuples, r_tuple)
    }
  }

  return relation
}


func relationsTypeMatches(lrelation, rrelation *unit.Relation) bool {
  return (lrelation.Tname == rrelation.Tname)
}

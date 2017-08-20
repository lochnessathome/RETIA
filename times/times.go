package times

import (
  "RETIA/unit"
  "RETIA/join"
)


func Create(lrelation, rrelation *unit.Relation, vname string) *unit.TimesStatement {
  if lrelation != nil && rrelation != nil {
    statement := new(unit.TimesStatement)

    statement.Lrelation = lrelation
    statement.Rrelation = rrelation

    statement.Vname = unit.FormatLetter(vname)

    return statement
  } else {
    return nil
  }
}


func Eval(statement *unit.TimesStatement) *unit.Relation {
  relation := new(unit.Relation)

  relation.Tname = statement.Lrelation.Tname
  relation.Vname = statement.Vname

  for _, l_tuple := range statement.Lrelation.Tuples {
    for _, r_tuple := range statement.Rrelation.Tuples {
      m_tuple := join.MergeTuples(l_tuple, r_tuple)
      relation.Tuples = append(relation.Tuples, m_tuple)
    }
  }

  return relation
}


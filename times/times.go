package times

import (
  "RETIA/unit"
  "RETIA/tuple"
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
      m_tuple := mergeTuples(l_tuple, r_tuple)
      relation.Tuples = append(relation.Tuples, m_tuple)
    }
  }

  return relation
}

func mergeTuples(l_tuple, r_tuple *unit.Tuple) *unit.Tuple {
  m_components := make([]*unit.Component, 0)

  for _, component := range l_tuple.Components {
    present := false

    for _, m_component := range m_components {
      if m_component.Aname == component.Aname && m_component.Atype == component.Atype {
        present = true
      }
    }

    if !present {
      m_components = append(m_components, component)
    }
  }

  for _, component := range r_tuple.Components {
    present := false

    for _, m_component := range m_components {
      if m_component.Aname == component.Aname && m_component.Atype == component.Atype {
        present = true
      }
    }

    if !present {
      m_components = append(m_components, component)
    }
  }

  return tuple.Create(m_components, "")
}


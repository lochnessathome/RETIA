package join

import (
  "RETIA/unit"
  "RETIA/tuple"
)


func Create(lrelation, rrelation *unit.Relation, vname string) *unit.JoinStatement {
  if lrelation != nil && rrelation != nil {
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

  c_attributes := findCommonAttributes(statement.Lrelation, statement.Rrelation)

  for _, l_tuple := range statement.Lrelation.Tuples {
    for _, r_tuple := range statement.Rrelation.Tuples {
      matches := false
      errors := false

      for _, c_attribute := range c_attributes {
        lcvalue := ""
        rcvalue := ""

        for _, l_component := range l_tuple.Components {
          if l_component.Aname == c_attribute.Aname && l_component.Atype == c_attribute.Atype {
            lcvalue = l_component.Cvalue
          }
        }

        for _, r_component := range r_tuple.Components {
          if r_component.Aname == c_attribute.Aname && r_component.Atype == c_attribute.Atype {
            rcvalue = r_component.Cvalue
          }
        }

        if lcvalue == rcvalue {
          matches = true
        } else {
          errors = true
        }
      }

      if matches && !errors {
        m_tuple := MergeTuples(l_tuple, r_tuple)
        relation.Tuples = append(relation.Tuples, m_tuple)
      }
    }
  }

  return relation
}


func MergeTuples(l_tuple, r_tuple *unit.Tuple) *unit.Tuple {
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

func findCommonAttributes(lrelation, rrelation *unit.Relation) []*unit.Component { 
  attributes := make([]*unit.Component, 0)

  ltuple := lrelation.Tuples[0]
  rtuple := rrelation.Tuples[0]

  for _, l_component := range ltuple.Components {
    for _, r_component := range rtuple.Components {
      if l_component.Aname == r_component.Aname && l_component.Atype == r_component.Atype {
        attribute := new(unit.Component)

        attribute.Aname = l_component.Aname
        attribute.Atype = l_component.Atype

        attributes = append(attributes, attribute)
        break
      }
    }
  }

  return attributes
}
 

package projection

import (
  "fmt"

  "RETIA/unit"
  "RETIA/tuple"
  "RETIA/relation"
)


func Create(relation *unit.Relation, anames []string) *unit.ProjectionStatement {
  if relation != nil && attributesValid(relation, anames) {
    statement := new(unit.ProjectionStatement)

    statement.Relation = relation
    statement.Anames = anames

    return statement
  } else {
    return nil
  }
}


func Eval(statement *unit.ProjectionStatement) *unit.Relation {
  if statement != nil {
    tuples := make([]*unit.Tuple, 0)

    for _, tuple := range statement.Relation.Tuples {
      tuples = append(tuples, buildTuple(tuple, statement.Anames))
    }

    return relation.Create(tuples)
  } else {
    return nil
  }
}


func attributesValid(relation *unit.Relation, anames []string) bool {
  tuple := relation.Tuples[0]

  for _, aname := range anames {
    present := false

    for _, component := range tuple.Components {
      if component.Aname == aname {
        present = true
      }
    }

    if !present {
      fmt.Printf("Attribute %s not found.\n", aname)
      return false
    }
  }

  return true
}

func buildTuple(orig_tuple *unit.Tuple, anames []string) *unit.Tuple {
  components := make([]*unit.Component, 0)

  for _, component := range orig_tuple.Components {
    for _, aname := range anames {
      if component.Aname == aname {
        components = append(components, component)

        break
      }
    }
  }

  return tuple.Create(components, "")  
}


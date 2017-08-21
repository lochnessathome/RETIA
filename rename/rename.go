package rename

import (
  "fmt"

  "RETIA/unit"
  "RETIA/component"
  "RETIA/tuple"
  "RETIA/relation"
)


func CreateExpression(laname, raname string) *unit.RenameExpression {
  expr := new(unit.RenameExpression)

  expr.Laname = unit.FormatLetter(laname)
  expr.Raname = unit.FormatLetter(raname)

  return expr
}

func Create(relation *unit.Relation, expressions []*unit.RenameExpression) *unit.RenameStatement {
  if relation != nil && expressionsValid(relation, expressions) {
    statement := new(unit.RenameStatement)

    statement.Relation = relation
    statement.Expressions = expressions

    return statement
  } else {
    return nil
  }
}


func Eval(statement *unit.RenameStatement) *unit.Relation {
  if statement != nil {
    tuples := make([]*unit.Tuple, 0)

    for _, tuple := range statement.Relation.Tuples {
      tuples = append(tuples, buildTuple(tuple, statement.Expressions))
    }

    return relation.Create(tuples)
  } else {
    return nil
  }
}


func expressionsValid(relation *unit.Relation, expressions []*unit.RenameExpression) bool {
  tuple := relation.Tuples[0]

  for _, expression := range expressions {
    present := false

    for _, component := range tuple.Components {
      if expression.Laname == component.Aname {
        present = true
      }
    }

    if !present {
      fmt.Printf("Attribute %s not found. \n", expression.Laname)

      return false
    }
  }

  for _, expression := range expressions {
    counter := 0

    for _, int_expression := range expressions {
      if expression.Raname == int_expression.Raname {
        counter = counter + 1
      }
    }

    if counter != 1 {
      fmt.Printf("Attribute %s not unique. \n", expression.Raname)

      return false
    }
  }

  return true
}

func buildTuple(orig_tuple *unit.Tuple, expressions []*unit.RenameExpression) *unit.Tuple {
  components := make([]*unit.Component, 0)

  for _, orig_component := range orig_tuple.Components {
    matches := false 

    for _, expression := range expressions {
      if orig_component.Aname == expression.Laname {
        n_component := component.Create(expression.Raname, orig_component.Atype, orig_component.Cvalue, orig_component.Atype) 
        components = append(components, n_component)

        matches = true
      }
    }

    if !matches {
      components = append(components, orig_component)
    }
  }

  return tuple.Create(components, "")  
}


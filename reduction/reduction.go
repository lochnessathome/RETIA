package reduction

import (
  "fmt"
  "strconv"

  "RETIA/unit"

  "RETIA/messages"
)


func Create(relation *unit.Relation, expr *unit.CompareExpression) *unit.ReductionStatement {
  if relation != nil && attributesValid(relation, expr) {
    statement := new(unit.ReductionStatement)

    statement.Relation = relation
    statement.CompareExpression = expr

    return statement
  } else {
    return nil
  }
}


func Eval(statement *unit.ReductionStatement) *unit.Relation {
  if statement != nil {
    relation := new(unit.Relation)

    relation.Tname = statement.Relation.Tname

    expr := statement.CompareExpression

    if len(expr.Raname) != 0 {
      for _, tuple := range statement.Relation.Tuples {
        if tupleAttrsMatches(tuple, expr.Operator, expr.Laname, expr.Raname) {
          relation.Tuples = append(relation.Tuples, tuple)
        }
      }
    } else {
      for _, tuple := range statement.Relation.Tuples {
        if tupleValuesMatches(tuple, expr.Operator, expr.Laname, expr.Rcvalue) {
          relation.Tuples = append(relation.Tuples, tuple)
        }
      }
    }

    return relation
  } else {
    return nil
  }
}


func attributesValid(relation *unit.Relation, expr *unit.CompareExpression) bool {
  loptype := ""
  roptype := ""

  _, loptype = findAttrByName(relation, expr.Laname)

  if len(loptype) == 0 {
    fmt.Printf("Given attribute %s not found in relation %s. \n", expr.Laname, relation.Vname)

    return false
  }

  if len(expr.Raname) != 0 {
    _, roptype = findAttrByName(relation, expr.Raname)

    if len(roptype) == 0 {
      fmt.Printf("Given attribute %s not found in relation %s. \n", expr.Raname, relation.Vname)

      return false
    }
  } else {
    roptype = expr.Rctype
  }

  if loptype != roptype {
    fmt.Printf("Can't compare values of different types: %s and %s. \n", loptype, roptype)

    return false
  }

  if !operationPermitted(expr.Operator, loptype) {
    fmt.Printf("Type %s doesn't support %s operator. \n", loptype, expr.Operator)

    return false
  }

  return true
}

func findAttrByName(relation *unit.Relation, aname string) (string, string) {
  for _, component := range relation.Tuples[0].Components {
    if component.Aname == aname {
      return component.Aname, component.Atype
    }
  }

  return "", ""
}

func operationPermitted(operator, atype string) bool {
  if atype == "integer" || atype == "rational" {
    switch operator {
      case "=", "<>", ">", ">=", "<", "<=":
        return true
    }
    return false
  }

  if atype == "char" || atype == "boolean" {
    switch operator {
      case "=", "<>":
        return true
    }
    return false
  }

  return false
}

func tupleAttrsMatches(tuple *unit.Tuple, operator, laname, raname string) bool {
  lvalue := ""
  rvalue := ""
  atype := ""

  for _, component := range tuple.Components {
    if component.Aname == laname {
      lvalue = component.Cvalue
      atype = component.Atype
    } else if component.Aname == raname {
      rvalue = component.Cvalue
    }
  }

  return compareValues(operator, lvalue, rvalue, atype)
}

func tupleValuesMatches(tuple *unit.Tuple, operator, laname, rcvalue string) bool {
  lvalue := ""
  atype := ""

  for _, component := range tuple.Components {
    if component.Aname == laname {
      lvalue = component.Cvalue
      atype = component.Atype
    }
  }

  return compareValues(operator, lvalue, rcvalue, atype)
}

func compareValues(operator, lvalue, rvalue, atype string) bool {
  if atype == "integer" {
    lvalue, _ := strconv.ParseInt(lvalue, 10, 64)
    rvalue, _ := strconv.ParseInt(rvalue, 10, 64)

    switch operator {
      case "=":
        return (lvalue == rvalue)
      case "<>":
        return (lvalue != rvalue)
      case ">":
        return (lvalue > rvalue)
      case ">=":
        return (lvalue >= rvalue)
      case "<":
        return (lvalue < rvalue)
      case "<=":
        return (lvalue <= rvalue)
    }

  } else if atype == "rational" {
    lvalue, _  := strconv.ParseFloat(lvalue, 64)
    rvalue, _  := strconv.ParseFloat(rvalue, 64)

    switch operator {
      case "=":
        return (lvalue == rvalue)
      case "<>":
        return (lvalue != rvalue)
      case ">":
        return (lvalue > rvalue)
      case ">=":
        return (lvalue >= rvalue)
      case "<":
        return (lvalue < rvalue)
      case "<=":
        return (lvalue <= rvalue)
    }

  } else {
    switch operator {
      case "=":
        return (lvalue == rvalue)
      case "<>":
        return (lvalue != rvalue)
    }
  }

  return false
}


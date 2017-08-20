package compare

import (
  "RETIA/unit"
)


func Create(operator, laname, raname, rcvalue, rctype string) *unit.CompareExpression {
  expr := new(unit.CompareExpression)

  expr.Operator = operator
  expr.Laname = unit.FormatLetter(laname)

  if len(raname) != 0 {
    expr.Raname = unit.FormatLetter(raname)
  } else {
    expr.Rcvalue = rcvalue
    expr.Rctype = rctype
  }

  return expr
}


package compare

import (
  "RETIA/unit"
)


func Create(operator, laname, raname, rcvalue, rctype string) *unit.Compare {
  compare := new(unit.Compare)

  compare.Operator = operator
  compare.Laname = unit.FormatLetter(laname)

  if len(raname) != 0 {
    compare.Raname = unit.FormatLetter(raname)
  } else {
    compare.Rcvalue = rcvalue
    compare.Rctype = rctype
  }

  return compare
}


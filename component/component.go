package component

import (
  "RETIA/unit"
)


func Create(aname, atype, cvalue string) *unit.Component {
  component := new(unit.Component)

  component.Aname = aname
  component.Atype = atype
  component.Cvalue = cvalue

  return component
}


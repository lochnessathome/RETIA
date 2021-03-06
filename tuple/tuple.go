package tuple

import (
  "sort"

  "RETIA/unit"
)


type ByAname []*unit.Component

func (a ByAname) Len() int           { return len(a) }
func (a ByAname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAname) Less(i, j int) bool { return a[i].Aname < a[j].Aname }


func Create(components []*unit.Component) *unit.Tuple {
  if len(components) != 0 && componentsValid(components) {
    tuple := new(unit.Tuple)

    sort.Sort(ByAname(components))
    for _, component := range components {
      tuple.Tname = tuple.Tname + unit.FormatTypeStr(component.Aname, component.Atype)
      tuple.Hash = tuple.Hash + component.Aname + component.Atype + component.Cvalue
    }

    tuple.Hash = unit.FormatHash(tuple.Hash)

    tuple.Components = components

    return tuple
  } else {
    return nil
  }
}


func componentsValid(components []*unit.Component) bool {
  for _, component := range components {
    if component == nil {
      return false
    }
  }

  return true
}


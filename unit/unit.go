package unit

type Component struct {
  Aname string
  Atype string
  Cvalue string
}

type Tuple struct {
  Vname string
  Tname string
  Components []*Component
}

type Relation struct {
  Vname string
  Tname string
  Tuples []*Tuple
}


package main

import (
  "os"

  "RETIA/unit"
  "RETIA/show"
)


type Session struct {
  tuples []*unit.Tuple
  relations []*unit.Relation
}


func (session *Session) Query(tuple *unit.Tuple, relation *unit.Relation, where *unit.Where) {
  if tuple != nil {
    if findRelation(session, tuple.Vname) == nil {

      foundTuple := findTuple(session, tuple.Vname)

      if foundTuple != nil {
        *foundTuple = *tuple
      } else {
        session.tuples = append(session.tuples, tuple)
      }

      show.Tuple(tuple)

    } else {
      show.VnameBusy(tuple.Vname)
    }
  }

  if relation != nil {
    if findTuple(session, relation.Vname) == nil {

      foundRelation := findRelation(session, relation.Vname) 
  
      if foundRelation != nil {
        *foundRelation = *relation
      } else {
        session.relations = append(session.relations, relation)
      }

      show.Relation(relation)
    } else {
      show.VnameBusy(relation.Vname)
    }
  }

  if where != nil {
    show.Where(where)

    nrel := new(unit.Relation)

    nrel.Tname = where.Relation.Tname
    nrel.Vname = where.Relation.Vname

    compare := where.Compare

    if len(compare.Raname) != 0 {
      for _, tuple := range where.Relation.Tuples {
        lvalue := ""
        rvalue := ""

        for _, component := range tuple.Components {
          if component.Aname == compare.Laname {
            lvalue = component.Cvalue
          }

          if component.Aname == compare.Raname {
            rvalue = component.Cvalue
          }

          if lvalue == rvalue {
            nrel.Tuples = append(nrel.Tuples, tuple)
          }
        }
      }
    } else {
      for _, tuple := range where.Relation.Tuples {
        for _, component := range tuple.Components {

          if component.Aname == compare.Laname {
            if component.Cvalue == compare.Rcvalue { 
              nrel.Tuples = append(nrel.Tuples, tuple)
            }
          }

        }
      }
    }

  // return nrel

  show.Relation(nrel)

  }
}

func (session *Session) Call(vname string) (*unit.Tuple, *unit.Relation) {
  vname = unit.FormatLetter(vname)

  foundTuple := findTuple(session, vname)

  if foundTuple != nil {
    return foundTuple, nil
  }

  foundRelation := findRelation(session, vname)

  if foundRelation != nil {
    return nil, foundRelation
  }

  show.VnameMissing(vname)

  return nil, nil
}


func findTuple(session *Session, vname string) *unit.Tuple {
  if len(vname) != 0 {
    for _, tuple := range session.tuples {
      if tuple.Vname == vname {
        return tuple
      }
    }
  }

  return nil
}

func findRelation(session *Session, vname string) *unit.Relation {
  if len(vname) != 0 {
    for _, relation := range session.relations {
      if relation.Vname == vname {
        return relation
      }
    }
  }

  return nil
}


func main() {
  yyErrorVerbose = true

  session := new(Session)

  for {
    yyParse(NewLexerWithInit(os.Stdin, func(y *Lexer) { y.p = session }))
  }
}


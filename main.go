package main

import (
  "os"

  "RETIA/unit"
  "RETIA/messages"
)


type Session struct {
  tuples []*unit.Tuple
  relations []*unit.Relation
}


func (session *Session) Query(tuple *unit.Tuple, relation *unit.Relation, assign string) {

  if len(assign) != 0 {

    if relation != nil {
      relation.Vname = unit.FormatLetter(assign)

      session.Query(nil, relation, "")
    }

  } else if relation != nil {

    if findTuple(session, relation.Vname) == nil {
      foundRelation := findRelation(session, relation.Vname)

      if foundRelation != nil {
        *foundRelation = *relation
      } else {
        session.relations = append(session.relations, relation)
      }

      messages.Relation(relation)

    } else {
      messages.VnameBusy(relation.Vname)
    }

  } else if tuple != nil {

    if findRelation(session, tuple.Vname) == nil {
      foundTuple := findTuple(session, tuple.Vname)

      if foundTuple != nil {
        *foundTuple = *tuple
      } else {
        session.tuples = append(session.tuples, tuple)
      }

      messages.Tuple(tuple)

    } else {
      messages.VnameBusy(tuple.Vname)
    }

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

  messages.VnameMissing(vname)

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


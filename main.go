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


func (session *Session) Query(tuple *unit.Tuple, relation *unit.Relation) {
  if tuple != nil {
    foundTuple := findTuple(session, tuple.Vname)

    if foundTuple != nil {
      *foundTuple = *tuple
    } else {
      session.tuples = append(session.tuples, tuple)
    }

    show.Tuple(tuple)
  }

  if relation != nil {
    foundRelation := findRelation(session, relation.Vname)
  
    if foundRelation != nil {
      *foundRelation = *relation
    } else {
      session.relations = append(session.relations, relation)
    }

    show.Relation(relation)
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

  return nil, nil
}


func findTuple(session *Session, vname string) *unit.Tuple {
  for _, tuple := range session.tuples {
    if tuple.Vname == vname {
      return tuple
    }
  }

  return nil
}

func findRelation(session *Session, vname string) *unit.Relation {
  for _, relation := range session.relations {
    if relation.Vname == vname {
      return relation
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


package main

import (
  "os"

  "RETIA/unit"
  "RETIA/show"
)


type Session struct {
  data string

  tuples []*unit.Tuple
  relations []*unit.Relation
}


func (session *Session) Query(tuple *unit.Tuple, relation *unit.Relation) {
  if tuple != nil {
    session.tuples = append(session.tuples, tuple)

    show.Tuple(tuple)
  }

  if relation != nil {
    session.relations = append(session.relations, relation)

    show.Relation(relation)
  }


}

func (session *Session) Call(vname string) (*unit.Tuple, *unit.Relation) {
  vname = unit.FormatLetter(vname)

  for _, tuple := range session.tuples {
    if tuple.Vname == vname {
      return tuple, nil
    }
  }

  for _, relation := range session.relations {
    if relation.Vname == vname {
      return nil, relation
    }
  }

  return nil, nil
}


func main() {
  yyErrorVerbose = true

  session := new(Session)

  for {
    yyParse(NewLexerWithInit(os.Stdin, func(y *Lexer) { y.p = session }))
  }
}


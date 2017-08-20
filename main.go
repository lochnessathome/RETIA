package main

import (
  "os"

  "RETIA/unit"
  "RETIA/show"

  "RETIA/reduction"
  "RETIA/union"
  "RETIA/intersection"
  "RETIA/minus"
  "RETIA/times"
  "RETIA/join"
)


type Session struct {
  tuples []*unit.Tuple
  relations []*unit.Relation
}


func (session *Session) Query(tuple *unit.Tuple, relation *unit.Relation, reduction_st *unit.ReductionStatement, union_st *unit.UnionStatement, intersection_st *unit.IntersectionStatement, minus_st *unit.MinusStatement, times_st *unit.TimesStatement, join_st *unit.JoinStatement, verbose bool) {
  if tuple != nil {
    if findRelation(session, tuple.Vname) == nil {

      foundTuple := findTuple(session, tuple.Vname)

      if foundTuple != nil {
        *foundTuple = *tuple
      } else {
        session.tuples = append(session.tuples, tuple)
      }

      if verbose {
        show.Tuple(tuple)
      }

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

      if verbose {
        show.Relation(relation)
      }

    } else {
      show.VnameBusy(relation.Vname)
    }
  }

  if reduction_st != nil {
    erelation := reduction.Eval(reduction_st)

    session.Query(nil, erelation, nil, nil, nil, nil, nil, nil, false)
  }

  if union_st != nil {
    erelation := union.Eval(union_st)

    session.Query(nil, erelation, nil, nil, nil, nil, nil, nil, false)
  }

  if intersection_st != nil {
    erelation := intersection.Eval(intersection_st)

    session.Query(nil, erelation, nil, nil, nil, nil, nil, nil, false)
  }

  if minus_st != nil {
    erelation := minus.Eval(minus_st)

    session.Query(nil, erelation, nil, nil, nil, nil, nil, nil, false)
  }

  if times_st != nil {
    erelation := times.Eval(times_st)

    session.Query(nil, erelation, nil, nil, nil, nil, nil, nil, false)
  }

  if join_st != nil {
    erelation := join.Eval(join_st)

    session.Query(nil, erelation, nil, nil, nil, nil, nil, nil, false)
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


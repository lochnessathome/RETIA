%{
  package main

  import (
    "RETIA/unit"
    "RETIA/component"
    "RETIA/tuple"
    "RETIA/relation"

    "RETIA/compare"
    "RETIA/reduction"

    "RETIA/union"
    "RETIA/intersection"
    "RETIA/minus"
    "RETIA/times"
    "RETIA/join"
  )
%}

%union {
  s string
  ctype string

  component *unit.Component
  components []*unit.Component

  tuple *unit.Tuple
  tuples []*unit.Tuple

  relation *unit.Relation

  compare_expr *unit.CompareExpression
  reduction_st *unit.ReductionStatement

  union_st *unit.UnionStatement
  intersection_st *unit.IntersectionStatement
  minus_st *unit.MinusStatement
  times_st *unit.TimesStatement
  join_st *unit.JoinStatement
}

%token ID
%token ASSIGN
%token RELATION
%token TUPLE
%token REDUCTION
%token UNION
%token INTERSECTION
%token MINUS
%token TIMES
%token JOIN
%token T_INTEGER
%token T_RATIONAL
%token T_CHAR
%token T_BOOLEAN
%token V_INTEGER
%token V_RATIONAL
%token V_CHAR
%token V_BOOLEAN
%token V_COMPARE

%%
input:				
			| input line		
			;

line:			'\n'
			| query '\n'						{ cast(yylex).Query($1.tuple, $1.relation, $1.reduction_st, $1.union_st, $1.intersection_st, $1.minus_st, $1.times_st, $1.join_st) }
			;

query:			tuple							{ $$.tuple = $1.tuple }
			| tuple_var						{ $$.tuple = $1.tuple }
			| relation						{ $$.relation = $1.relation }
                        | relation_var                                  	{ $$.relation = $1.relation }
			| reduction						{ $$.reduction_st = $1.reduction_st }
                        | reduction_var                                     	{ $$.reduction_st = $1.reduction_st }
                        | union                                             	{ $$.union_st = $1.union_st }
                        | union_var                                         	{ $$.union_st = $1.union_st }
                        | intersection                                          { $$.intersection_st = $1.intersection_st }
                        | intersection_var                                      { $$.intersection_st = $1.intersection_st }
                        | minus                                                 { $$.minus_st = $1.minus_st }
                        | minus_var                                             { $$.minus_st = $1.minus_st }
                        | times                                                 { $$.times_st = $1.times_st }
                        | times_var                                             { $$.times_st = $1.times_st }
                        | join                                                  { $$.join_st = $1.join_st }
                        | join_var                                              { $$.join_st = $1.join_st }
			;

relation:		RELATION '{' tuples_commalist '}'			{ $$.relation = relation.Create($3.tuples, "") }
			| ID 							{ $$.tuple, $$.relation = cast(yylex).Call($1.s) }
			;

relation_var:		ID ASSIGN RELATION '{' tuples_commalist '}'		{ $$.relation = relation.Create($5.tuples, $1.s) }
			;

tuples_commalist:   	tuple	                                        	{ $$.tuples = append($$.tuples, $1.tuple) }
                        | tuples_commalist ',' tuple		        	{ $$.tuples = append($$.tuples, $3.tuple) }
                        ;

tuple:             	TUPLE '{' components_commalist '}'			{ $$.tuple = tuple.Create($3.components, "") }
			| ID							{ $$.tuple, $$.relation = cast(yylex).Call($1.s) }
			;

tuple_var:		ID ASSIGN TUPLE '{' components_commalist '}'		{ $$.tuple = tuple.Create($5.components, $1.s) }
                        ;


reduction:		relation REDUCTION '(' compare_expression ')'		{ $$.reduction_st = reduction.Create($1.relation, $4.compare_expr, "") }
			;

reduction_var:          ID ASSIGN relation REDUCTION '(' compare_expression ')' { $$.reduction_st = reduction.Create($3.relation, $6.compare_expr, $1.s) }
                        ;

union:		        relation UNION relation           			{ $$.union_st = union.Create($1.relation, $3.relation, "") }
                        ;

union_var:          	ID ASSIGN relation UNION relation                       { $$.union_st = union.Create($3.relation, $5.relation, $1.s) }
                        ;

intersection:           relation INTERSECTION relation                          { $$.intersection_st = intersection.Create($1.relation, $3.relation, "") }
                        ;

intersection_var:       ID ASSIGN relation INTERSECTION relation                { $$.intersection_st = intersection.Create($3.relation, $5.relation, $1.s) }
                        ;

minus:                  relation MINUS relation                                 { $$.minus_st = minus.Create($1.relation, $3.relation, "") }
                        ;

minus_var:              ID ASSIGN relation MINUS relation                       { $$.minus_st = minus.Create($3.relation, $5.relation, $1.s) }
                        ;

times:                  relation TIMES relation                                 { $$.times_st = times.Create($1.relation, $3.relation, "") }
                        ;

times_var:              ID ASSIGN relation TIMES relation                       { $$.times_st = times.Create($3.relation, $5.relation, $1.s) }
                        ;

join:                   relation JOIN relation                                  { $$.join_st = join.Create($1.relation, $3.relation, "") }
                        ;

join_var:               ID ASSIGN relation JOIN relation                        { $$.join_st = join.Create($3.relation, $5.relation, $1.s) }
                        ;


compare_expression:	attribute_name V_COMPARE attribute_name			{ $$.compare_expr = compare.Create($2.s, $1.s, $3.s, "", "") }
			| attribute_name V_COMPARE component_value		{ $$.compare_expr = compare.Create($2.s, $1.s, "", $3.s, $3.ctype) }
			;

components_commalist:   component						{ $$.components = append($$.components, $1.component) }
                        | components_commalist ',' component			{ $$.components = append($$.components, $3.component) }
                        ;

component:		attribute_name attribute_type component_value		{ $$.component = component.Create($1.s, $2.s, $3.s, $3.ctype) }
                        ;

attribute_name:		ID
			;

attribute_type:		built_in_type
			;

component_value:	V_INTEGER						{ $$.ctype = "integer" }
			| V_RATIONAL						{ $$.ctype = "rational" }
			| V_CHAR						{ $$.ctype = "char" }
			| V_BOOLEAN						{ $$.ctype = "boolean" }
			;

built_in_type:		T_INTEGER
			| T_RATIONAL
			| T_CHAR
			| T_BOOLEAN
			;

%%

func cast(y yyLexer) *Session { return y.(*Lexer).p }


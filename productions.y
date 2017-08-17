%{
  package main

  import (
    "RETIA/unit"
    "RETIA/component"
    "RETIA/tuple"
    "RETIA/relation"
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
}

%token ID
%token ASSIGN
%token RELATION
%token TUPLE
%token T_INTEGER
%token T_RATIONAL
%token T_CHAR
%token T_BOOLEAN
%token V_INTEGER
%token V_RATIONAL
%token V_CHAR
%token V_BOOLEAN

%%
input:				
			| input line		
			;

line:			'\n'
			| query '\n'					{ cast(yylex).Query($1.tuple, $1.relation) }
			;

query:			tuple						{ $$.tuple = $1.tuple }
			| tuple_var					{ $$.tuple = $1.tuple }
			| relation					{ $$.relation = $1.relation }
                        | relation_var                                  { $$.relation = $1.relation }
			| var_call					{ $$.relation = $1.relation; $$.tuple = $1.tuple }						
			;

var_call:		ID						{ $$.tuple, $$.relation = cast(yylex).Call($1.s) }
			;

relation:		RELATION '{' tuples_commalist '}'		{ $$.relation = relation.Create($3.tuples, "") }
			;

relation_var:		ID ASSIGN RELATION '{' tuples_commalist '}'	{ $$.relation = relation.Create($5.tuples, $1.s) }
			;

tuples_commalist:   	tuple	                                        { $$.tuples = append($$.tuples, $1.tuple) }
                        | tuples_commalist ',' tuple		        { $$.tuples = append($$.tuples, $3.tuple) }
                        ;

tuple:             	TUPLE '{' components_commalist '}'		{ $$.tuple = tuple.Create($3.components, "") }
			;

tuple_var:		ID ASSIGN TUPLE '{' components_commalist '}'	{ $$.tuple = tuple.Create($5.components, $1.s) }
                        ;

components_commalist:   component					{ $$.components = append($$.components, $1.component) }
                        | components_commalist ',' component		{ $$.components = append($$.components, $3.component) }
                        ;

component:		attribute_name attribute_type component_value	{ $$.component = component.Create($1.s, $2.s, $3.s, $3.ctype) }
                        ;

attribute_name:		ID
			;

attribute_type:		built_in_type
			;

component_value:	V_INTEGER					{ $$.ctype = "integer" }
			| V_RATIONAL					{ $$.ctype = "rational" }
			| V_CHAR					{ $$.ctype = "char" }
			| V_BOOLEAN					{ $$.ctype = "boolean" }
			;

built_in_type:		T_INTEGER
			| T_RATIONAL
			| T_CHAR
			| T_BOOLEAN
			;

%%

func cast(y yyLexer) *Session { return y.(*Lexer).p }


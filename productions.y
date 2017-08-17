%{
  package main

  import (
    "RETIA/unit"
    "RETIA/component"
    "RETIA/tuple"
  )
%}

%union {
  s string
  ctype string

  component *unit.Component
  components []*unit.Component

  tuple *unit.Tuple
}

%token ID
%token ASSIGN
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
			| query '\n'					{ cast(yylex).Query($1.tuple) }
			;

query:			tuple						{ $$.tuple = $1.tuple }
			;


tuple:             	TUPLE '{' components_commalist '}'		{ $$.tuple = tuple.Create($3.components, "") }
			| ID ASSIGN TUPLE '{' components_commalist '}'	{ $$.tuple = tuple.Create($5.components, $1.s) }
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

func cast(y yyLexer) *Result { return y.(*Lexer).p }


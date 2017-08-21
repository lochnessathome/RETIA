%{
  package main

  import (
    "RETIA/unit"
    "RETIA/component"
    "RETIA/tuple"
    "RETIA/relation"

    "RETIA/union"
    "RETIA/intersection"
    "RETIA/minus"
    "RETIA/times"
    "RETIA/join"

    "RETIA/reduction"
    "RETIA/compare"

    "RETIA/projection"

    "RETIA/rename"
  )
%}

%union {
  s string
  ctype string

  assign string

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

  attribute_names_list []string
  projection_st *unit.ProjectionStatement

  rename_expr *unit.RenameExpression
  rename_expr_list []*unit.RenameExpression
  rename_st *unit.RenameStatement
}

%token ID
%token ASSIGN
%token RELATION
%token TUPLE
%token AS
%token PROJECTION
%token RENAME
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

%left PROJECTION RENAME
%left REDUCTION
%left UNION INTERSECTION MINUS TIMES JOIN

%%
input:				
				| input line		
				;

line:				'\n'
				| query '\n'						{ cast(yylex).Query($1.tuple, $1.relation, $1.assign) }
				;

query:				tuple							{ $$.tuple = $1.tuple }
				| tuple_var						{ $$.tuple = $1.tuple }
				| relation						{ $$.relation = $1.relation }
                        	| ID ASSIGN relation	                                { $$.assign = $1.s; $$.relation = $3.relation }
				;

relation:			RELATION '{' tuples_commalist '}'			{ $$.relation = relation.Create($3.tuples) }
				| ID 							{ $$.tuple, $$.relation = cast(yylex).Call($1.s) }
				| union							{ $$.relation = union.Eval($1.union_st) }
				| intersection						{ $$.relation = intersection.Eval($1.intersection_st) }
                        	| minus                                                 { $$.relation = minus.Eval($1.minus_st) }
                        	| times                                                 { $$.relation = times.Eval($1.times_st) }
                        	| join                                                  { $$.relation = join.Eval($1.join_st) }
				| reduction						{ $$.relation = reduction.Eval($1.reduction_st) }
				| projection						{ $$.relation = projection.Eval($1.projection_st) }
                        	| rename                                                { $$.relation = rename.Eval($1.rename_st) }
				;

tuples_commalist:   		tuple	                                        	{ $$.tuples = append($$.tuples, $1.tuple) }
                        	| tuples_commalist ',' tuple		        	{ $$.tuples = append($$.tuples, $3.tuple) }
                        	;

tuple:             		TUPLE '{' components_commalist '}'			{ $$.tuple = tuple.Create($3.components, "") }
				| ID							{ $$.tuple, $$.relation = cast(yylex).Call($1.s) }
				;

tuple_var:			ID ASSIGN TUPLE '{' components_commalist '}'		{ $$.tuple = tuple.Create($5.components, $1.s) }
                        	;

union:		        	relation UNION relation           			{ $$.union_st = union.Create($1.relation, $3.relation) }
                        	;

intersection:           	relation INTERSECTION relation                          { $$.intersection_st = intersection.Create($1.relation, $3.relation) }
				;

minus:                  	relation MINUS relation                                 { $$.minus_st = minus.Create($1.relation, $3.relation) }
				;

times:                  	relation TIMES relation                                 { $$.times_st = times.Create($1.relation, $3.relation) }
				;

join:                   	relation JOIN relation                                  { $$.join_st = join.Create($1.relation, $3.relation) }
				;

reduction:			relation REDUCTION '(' compare_expression ')'		{ $$.reduction_st = reduction.Create($1.relation, $4.compare_expr) }
				;

compare_expression:		attribute_name V_COMPARE attribute_name			{ $$.compare_expr = compare.Create($2.s, $1.s, $3.s, "", "") }
				| attribute_name V_COMPARE component_value		{ $$.compare_expr = compare.Create($2.s, $1.s, "", $3.s, $3.ctype) }
				;

projection:			relation PROJECTION '(' attribute_names_commalist ')'	{ $$.projection_st = projection.Create($1.relation, $4.attribute_names_list) }
				;

attribute_names_commalist:	attribute_name						{ $$.attribute_names_list = append($$.attribute_names_list, $1.s) }
				| attribute_names_commalist ',' attribute_name		{ $$.attribute_names_list = append($$.attribute_names_list, $3.s) }
				;

rename: 	        	relation RENAME '(' rename_expressions_commalist ')'   	{ $$.rename_st = rename.Create($1.relation, $4.rename_expr_list) }
                        	;

rename_expressions_commalist:	rename_expression					{ $$.rename_expr_list = append($$.rename_expr_list, $1.rename_expr) }
				| rename_expressions_commalist ',' rename_expression	{ $$.rename_expr_list = append($$.rename_expr_list, $3.rename_expr) }
				;

rename_expression:		attribute_name AS attribute_name			{ $$.rename_expr = rename.CreateExpression($1.s, $3.s) }
				;

components_commalist:   	component						{ $$.components = append($$.components, $1.component) }
                        	| components_commalist ',' component			{ $$.components = append($$.components, $3.component) }
                        	;

component:			attribute_name attribute_type component_value		{ $$.component = component.Create($1.s, $2.s, $3.s, $3.ctype) }
                        	;

attribute_name:			ID
				;

attribute_type:			built_in_type
				;

component_value:		V_INTEGER						{ $$.ctype = "integer" }
				| V_RATIONAL						{ $$.ctype = "rational" }
				| V_CHAR						{ $$.ctype = "char" }
				| V_BOOLEAN						{ $$.ctype = "boolean" }
				;

built_in_type:			T_INTEGER
				| T_RATIONAL
				| T_CHAR
				| T_BOOLEAN
				;

%%

func cast(y yyLexer) *Session { return y.(*Lexer).p }


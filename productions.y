%{
  package main
  import "fmt"
%}

%union {
  s string

  attribute *Attribute
  attributes []*Attribute

  component *Component
  components []*Component

  tuple *Tuple
}

%token ID
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

query:			tuple_heading					{ $$.tuple = $1.tuple }
			| tuple_body					{ $$.tuple = $1.tuple }
			;

tuple_heading:		TUPLE '{' attributes_commalist '}'		{ $$.tuple = NewTupleHeading($3.attributes) }
			;

tuple_body:             TUPLE '{' components_commalist '}'		{ $$.tuple = NewTupleBody($3.components) }
                        ;


attributes_commalist:	attribute					{ $$.attributes = append($$.attributes, $1.attribute) }
			| attributes_commalist ',' attribute		{ $$.attributes = append($$.attributes, $3.attribute) }
			;

components_commalist:   component					{ $$.components = append($$.components, $1.component) }
                        | components_commalist ',' component		{ $$.components = append($$.components, $3.component) }
                        ;

attribute:	        attribute_name attribute_type			{ $$.attribute = NewAttribute($1.s, $2.s) }
                        ;

component:		attribute_name attribute_type attribute_value	{ $$.component = NewComponent($1.s, $2.s, $3.s) }
                        ;

attribute_name:		ID
			;

attribute_type:		built_in_type
			;

attribute_value:	V_INTEGER
			| V_RATIONAL
			| V_CHAR
			| V_BOOLEAN
			;

built_in_type:		T_INTEGER
			| T_RATIONAL
			| T_CHAR
			| T_BOOLEAN
			;

%%

func cast(y yyLexer) *Result { return y.(*Lexer).p }


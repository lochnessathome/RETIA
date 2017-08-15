%{
  package main
  import "fmt"
%}

%union {
  s string

  attribute *Attribute
  attributes []*Attribute

  tuple *Tuple
}

%token ID
%token NUMBER
%token TUPLE
%token INTEGER
%token RATIONAL
%token CHAR
%token BOOLEAN

%%
input:				
			| input line		
			;

line:			'\n'
			| query '\n'					{ cast(yylex).Query($1.tuple) }
			;

query:			tuple_heading					{ $$.tuple = $1.tuple }
			| tuple_body					{ }
			;

tuple_heading:		TUPLE '{' attributes_commalist '}'		{ $$.tuple = NewTuple($3.attributes) }
			;

tuple_body:             TUPLE '{' components_commalist '}'
                        ;


attributes_commalist:	attribute					{ $$.attributes = append($$.attributes, $1.attribute) }
			| attributes_commalist ',' attribute		{ $$.attributes = append($$.attributes, $3.attribute) }
			;

components_commalist:   component				
                        | components_commalist ',' component	
                        ;

attribute:	        attribute_name attribute_type			{ $$.attribute = NewAttribute($1.s, $2.s) }
                        ;

component:		attribute_name attribute_type attribute_value
                        ;

attribute_name:		ID
			;

attribute_type:		built_in_type
			;

attribute_value:	NUMBER						{ fmt.Printf("attribute value %s \n", $1.s) }

built_in_type:		INTEGER
			| RATIONAL
			| CHAR
			| BOOLEAN
			;

%%

func cast(y yyLexer) *Result { return y.(*Lexer).p }


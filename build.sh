#!/bin/bash

/home/dmitriy/homework/bin/nex -e=true terminal-def.nex

go tool yacc productions.y

printf '/NEX_END_OF_LEXER_STRUCT/i\np *Session\n.\nw\nq\n' | ed -s terminal-def.nn.go

go build main.go terminal-def.nn.go y.go


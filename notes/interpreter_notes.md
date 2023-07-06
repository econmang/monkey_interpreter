# Interpreters

## Interpreter Structure:
The interpreter will have 5 major parts:
- lexer
- parser
- abstract syntax tree (AST)
- internal object system
- evaluator

## Lexical Analysis:

Transforming source code to tokens is called lexical analysis ("lexing")
Tokens: Small, easily categorized data structs that are fed to the parser.
The parser takes the tokens and creates an Abstract Syntax Tree.

Lexing Example:
"let x = 5 + 5;" == Lexer ==> [
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]



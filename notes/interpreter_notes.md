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

## Parsing

A parser takes input and builds a data structure (often a parse tree or AST (Abstract Syntax Tree)), normally preceded by a lexical analyzer.
An example for JS parsing text input and outputting JS object:
```
> var input = '{"name": "Thorsten", "age": 28}';
> var output = JSON.parse(input)
> output
{ name: 'Thorsten', age: 28 }
> output.name
'Thorsten'
> output.age
28
```

Notes: 
    - An Abstract Syntax Tree vs a Syntax Tree -- the "abstract" implies that there are some details visible in source omitted in the AST.
    - There is no universal impl or output for an AST. Most impls are very similar, but there are diffs in the details

Stated succinctly, parsers take input (as text or tokens) and produce a data struct which represents the source code.

// token/token.go

package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"
    // Idents & Literals
    IDENT   = "IDENT"
    INT     = "INT"
    // Operators
    ASSIGN  = "="
    EQ      = "=="
    NOT_EQ  = "!="
    PLUS    = "+"
    MINUS   = "-"
    BANG    = "!"
    ASTERISK= "*"
    SLASH   = "/"
    LT      = "<"
    GT      = ">"
    // Delimiters
    COMMA       = ","
    SEMICOLON   = ";"
    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"
    // Keywords
    FUNCTION    = "FUNCTION"
    LET         = "LET"
    IF          = "IF"
    ELSE        = "ELSE"
    TRUE        = "TRUE"
    FALSE       = "FALSE"
    RETURN      = "RETURN"
)

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "if": IF,
    "else": ELSE,
    "true": TRUE,
    "false": FALSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

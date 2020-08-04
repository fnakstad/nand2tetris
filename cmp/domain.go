package main

type TokenType string
type Keyword string
type Symbol rune

type Token struct {
	Type       TokenType
	Keyword    Keyword
	Symbol     Symbol
	Identifier string
	IntVal     int
	StringVal  string
}

const (
	TokenTypeUnknown     TokenType = ""
	TokenTypeKeyword               = "keyword"
	TokenTypeSymbol                = "symbol"
	TokenTypeIdentifier            = "identifier"
	TokenTypeIntConst              = "integerConstant"
	TokenTypeStringConst           = "stringConstant"

	KeywordUnknown     Keyword = ""
	KeywordClass               = "class"
	KeywordMethod              = "method"
	KeywordFunction            = "function"
	KeywordConstructor         = "constructor"
	KeywordInt                 = "int"
	KeywordBoolean             = "boolean"
	KeywordChar                = "char"
	KeywordVoid                = "void"
	KeywordVar                 = "var"
	KeywordStatic              = "static"
	KeywordField               = "field"
	KeywordLet                 = "let"
	KeywordDo                  = "do"
	KeywordIf                  = "if"
	KeywordElse                = "else"
	KeywordWhile               = "while"
	KeywordReturn              = "return"
	KeywordTrue                = "true"
	KeywordFalse               = "false"
	KeywordNull                = "null"
	KeywordThis                = "this"

	SymbolUnknown          Symbol = 0
	SymbolLeftCurlyBrace          = '{'
	SymbolRightCurlyBrace         = '}'
	SymbolLeftParenthesis         = '('
	SymbolRightParenthesis        = ')'
	SymbolLeftSquareBrace         = '['
	SymbolRightSquareBrace        = ']'
	SymbolDot                     = '.'
	SymbolComma                   = ','
	SymbolSemiColon               = ';'
	SymbolPlus                    = '+'
	SymbolMinus                   = '-'
	SymbolAsterisk                = '*'
	SymbolSlash                   = '/'
	SymbolAnd                     = '&'
	SymbolOr                      = '|'
	SymbolLessThan                = '<'
	SymbolMoreThan                = '>'
	SymbolEqual                   = '='
)

var (
	Keywords = []Keyword{
		KeywordClass,
		KeywordMethod,
		KeywordFunction,
		KeywordConstructor,
		KeywordInt,
		KeywordBoolean,
		KeywordChar,
		KeywordVoid,
		KeywordVar,
		KeywordStatic,
		KeywordField,
		KeywordLet,
		KeywordDo,
		KeywordIf,
		KeywordElse,
		KeywordWhile,
		KeywordReturn,
		KeywordTrue,
		KeywordFalse,
		KeywordNull,
		KeywordThis,
	}
	SymbolsMap = map[rune]struct{}{
		SymbolLeftCurlyBrace:   struct{}{},
		SymbolRightCurlyBrace:  struct{}{},
		SymbolLeftParenthesis:  struct{}{},
		SymbolRightParenthesis: struct{}{},
		SymbolLeftSquareBrace:  struct{}{},
		SymbolRightSquareBrace: struct{}{},
		SymbolDot:              struct{}{},
		SymbolComma:            struct{}{},
		SymbolSemiColon:        struct{}{},
		SymbolPlus:             struct{}{},
		SymbolMinus:            struct{}{},
		SymbolAsterisk:         struct{}{},
		SymbolSlash:            struct{}{},
		SymbolAnd:              struct{}{},
		SymbolOr:               struct{}{},
		SymbolLessThan:         struct{}{},
		SymbolMoreThan:         struct{}{},
		SymbolEqual:            struct{}{},
	}
)

package main

type TokenType string
type Keyword string

const (
	TokenTypeUnknown     TokenType = ""
	TokenTypeKeyword               = "KEYWORD"
	TokenTypeSymbol                = "SYMBOL"
	TokenTypeIdentifier            = "IDENTIFIER"
	TokenTypeIntConst              = "INT_CONST"
	TokenTypeStringConst           = "STRING_CONST"

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
)

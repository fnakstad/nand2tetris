package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	lineCommentRegexp  = regexp.MustCompile(`^\/\/.*\n$`)
	blockCommentRegexp = regexp.MustCompile(`^\/\*(.*\n*)*\*\/$`)
	stringRegexp       = regexp.MustCompile(`^\".*\"$`)
	intRegexp          = regexp.MustCompile(`^[0-9]+$`)
	idRegexp           = regexp.MustCompile(`^[A-Za-z_]+[A-Za-z0-9_]*$`)
)

type Tokenizer struct {
	scanner *bufio.Scanner
	err     error
	text    string

	token Token
}

func NewTokenizer(r io.Reader) *Tokenizer {
	s := bufio.NewScanner(r)
	s.Split(ScanJackTokens)
	return &Tokenizer{
		scanner: s,
	}
}

func ScanJackTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) <= 0 {
		return 0, nil, nil
	}

	// Skip whitespace
	start := 0
	for start < len(data) && unicode.IsSpace(rune(data[start])) {
		start++
	}

	if len(data[start:]) <= 0 { // Need more data...
		return 0, nil, nil
	}

	firstRune := rune(data[start])
	if firstRune == '/' { // Either comment or '/' symbol
		if len(data[start:]) < 2 {
			return 0, nil, nil
		}
		secondRune := rune(data[start+1])

		if secondRune == '/' { // single-line comment
			if end := strings.Index(string(data[start:]), "\n"); end != -1 {
				return start + end + 1, data[start : start+end+1], nil
			}
			return 0, nil, nil
		}

		if secondRune == '*' { // block comment
			if end := strings.Index(string(data[start:]), "*/"); end != -1 {
				return start + end + 2, data[start : start+end+2], nil
			}
			return 0, nil, nil
		}

		// So, should be a '/' symbol. Just let fall through
	}

	// Check for symbols
	if isSymbol(firstRune) {
		return start + 1, []byte{data[start]}, nil
	}

	// Check for string constants
	if firstRune == '"' {
		if end := strings.Index(string(data[start+1:]), "\""); end != -1 {
			return start + end + 2, data[start : start+end+2], nil
		}
		return 0, nil, nil
	}

	// Otherwise just split on next space/symbol (keywords + identifiers)
	if end := strings.IndexFunc(string(data[start:]), func(r rune) bool {
		return unicode.IsSpace(r) || isSymbol(r)
	}); end != -1 {
		return start + end, data[start : start+end], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	// Not enough data?
	return 0, nil, nil
}

func (t *Tokenizer) Token() Token {
	return t.token
}

func (t *Tokenizer) Err() error {
	return t.err
}

func (t *Tokenizer) resetProps() {
	t.token = Token{}
	t.text = ""
}

func (t *Tokenizer) Next() bool {
	t.resetProps()

	skip := true
	for skip {
		if done := t.scanner.Scan(); !done {
			t.err = t.scanner.Err()
			return false
		}

		t.text = t.scanner.Text()

		// Skip comments
		if !blockCommentRegexp.MatchString(t.text) && !lineCommentRegexp.MatchString(t.text) {
			skip = false
		}
	}

	switch {
	case isSymbol(rune(t.text[0])):
		t.token.Type = TokenTypeSymbol
		t.token.Symbol = Symbol(t.text[0])
	case isKeyword(t.text):
		t.token.Type = TokenTypeKeyword
		t.token.Keyword = Keyword(t.text)
	case stringRegexp.MatchString(t.text):
		t.token.Type = TokenTypeStringConst
		t.token.StringVal = t.text[1 : len(t.text)-1]
	case intRegexp.MatchString(t.text):
		val, err := strconv.Atoi(t.text)
		if err != nil {
			t.err = errors.New("couldn't convert int const")
		}
		t.token.Type = TokenTypeIntConst
		t.token.IntVal = val
	case idRegexp.MatchString(t.text):
		t.token.Type = TokenTypeIdentifier
		t.token.Identifier = t.text
	default:
		t.err = fmt.Errorf("invalid token: %s", t.text)
		return false
	}

	return true
}

func isSymbol(r rune) bool {
	_, ok := SymbolsMap[r]
	return ok
}

func isKeyword(s string) bool {
	_, ok := KeywordsMap[s]
	return ok
}

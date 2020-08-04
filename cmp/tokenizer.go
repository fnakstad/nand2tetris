package main

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	commentPrefix = "//"
)

var (
	blockCommentRegexp = regexp.MustCompile(`/\*(.*?)\*/`)
	intRegexp          = regexp.MustCompile(`^[0-9]+`)
	idRegexp           = regexp.MustCompile(`^[A-Za-z_]+[A-Za-z0-9_]*`)
)

type Tokenizer struct {
	scanner *bufio.Scanner
	err     error
	text    string

	token Token
}

func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{
		scanner: bufio.NewScanner(r),
	}
}

func (t *Tokenizer) Token() Token {
	return t.token
}

func (t *Tokenizer) Err() error {
	return t.err
}

func (t *Tokenizer) resetProps() {
	t.token = Token{}
}

func isSymbol(r rune) bool {
	_, ok := SymbolsMap[r]
	return ok
}

func startsWithKeyword(w string) (bool, Keyword) {
	for _, kw := range Keywords {
		// TODO: identifiers starting with keywords...
		if strings.HasPrefix(w, string(kw)) {
			return true, kw
		}
	}
	return false, KeywordUnknown
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' //|| r == '\n'
}

func (t *Tokenizer) Next() bool {
	t.resetProps()

	for t.text == "" {
		if done := t.scanner.Scan(); !done {
			t.err = t.scanner.Err()
			return false
		}

		t.text = t.scanner.Text()
		t.text = stripCommentsAndWhitespace(t.text)
	}

	// Scan line for tokens
	advance := 0
	if isSymbol(rune(t.text[0])) {
		t.token.Type = TokenTypeSymbol
		t.token.Symbol = Symbol(t.text[0])
		advance = 1
	} else if ok, kw := startsWithKeyword(t.text); ok {
		t.token.Type = TokenTypeKeyword
		t.token.Keyword = kw
		advance = len(kw)
	} else if rune(t.text[0]) == '"' {
		i := strings.Index(t.text[1:], "\"")
		if i == -1 {
			t.err = errors.New("invalid token")
			return false
		}
		val := t.text[1 : i+1] /* Strip double quotes */
		t.token.Type = TokenTypeStringConst
		t.token.StringVal = val
		advance = i + 2
	} else if loc := intRegexp.FindStringIndex(t.text); loc != nil {
		t.token.Type = TokenTypeIntConst
		val, err := strconv.Atoi(t.text[loc[0]:loc[1]])
		if err != nil {
			t.err = errors.New("couldn't convert int const")
		}
		t.token.IntVal = val
		advance = loc[1]
	} else if loc := idRegexp.FindStringIndex(t.text); loc != nil {
		val := t.text[loc[0]:loc[1]]
		t.token.Type = TokenTypeIdentifier
		t.token.Identifier = val
		advance = loc[1]
	} else {
		t.err = errors.New("invalid token")
		return false
	}
	t.text = strings.TrimSpace(t.text[advance:])

	return true
}

func stripCommentsAndWhitespace(text string) string {
	if i := strings.Index(text, commentPrefix); i >= 0 {
		text = text[:i]
	}

	text = blockCommentRegexp.ReplaceAllString(text, "")

	return strings.TrimSpace(text)
}

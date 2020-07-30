package main

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

const (
	commentPrefix = "//"
)

var (
	blockCommentRegexp = regexp.MustCompile(`/\*(.*?)\*/`)
)

type Tokenizer struct {
	scanner    *bufio.Scanner
	err        error
	input      string
	tokenType  TokenType
	keyword    Keyword
	symbol     rune
	identifier string
	intVal     int
	stringVal  string
}

func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{
		scanner: bufio.NewScanner(r),
	}
}

func (t *Tokenizer) Err() error {
	return t.err
}

func (t *Tokenizer) Input() string {
	return t.input
}

func (t *Tokenizer) resetProps() {
	t.tokenType = TokenTypeUnknown
	t.keyword = KeywordUnknown
	t.symbol = 0
	t.identifier = ""
	t.intVal = 0
	t.stringVal = ""
}

func (t *Tokenizer) Next() bool {
	t.resetProps()

	var text string
	for text == "" {
		if done := t.scanner.Scan(); !done {
			t.err = t.scanner.Err()
			return false
		}

		text = t.scanner.Text()
		text = stripCommentsAndWhitespace(text)
	}

	t.input = text

	return true
}

func stripCommentsAndWhitespace(text string) string {
	if i := strings.Index(text, commentPrefix); i >= 0 {
		text = text[:i]
	}

	text = blockCommentRegexp.ReplaceAllString(text, "")

	return strings.TrimSpace(text)
}

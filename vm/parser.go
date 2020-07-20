package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	commentPrefix = "//"
)

type Parser struct {
	scanner *bufio.Scanner
	err     error
	cmdType CommandType
	arg1    string
	arg2    int
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		scanner: bufio.NewScanner(r),
	}
}

func (p *Parser) Err() error {
	return p.err
}

func (p *Parser) CommandType() CommandType {
	return p.cmdType
}

func (p *Parser) Arg1() string {
	return p.arg1
}

func (p *Parser) Arg2() int {
	return p.arg2
}

func (p *Parser) resetProps() {
	p.cmdType = CommandTypeUnknown
	p.arg1 = ""
	p.arg2 = 0
}

func (p *Parser) Parse() bool {
	p.resetProps()

	var text string
	for text == "" {
		if done := p.scanner.Scan(); !done {
			p.err = p.scanner.Err()
			return false
		}

		text = p.scanner.Text()
		text = stripCommentsAndWhitespace(text)
	}

	tokens := strings.Split(text, " ")
	if len(tokens) <= 0 {
		p.err = fmt.Errorf("couldn't parse: %s", text)
		return false
	}

	cmdType, ok := tokenCommandMap[tokens[0]]
	if !ok {
		p.err = fmt.Errorf("unrecognized token: %s", tokens[0])
		return false
	}
	p.cmdType = cmdType

	if _, ok := arithmeticCommand[p.CommandType()]; ok {
		p.arg1 = tokens[0]
	}
	switch cmdType {
	case CommandTypePush, CommandTypePop:
		if len(tokens) < 3 {
			p.err = fmt.Errorf("not enough args: %s", text)
			return false
		}

		p.arg1 = tokens[1]

		arg2, err := strconv.Atoi(tokens[2])
		if err != nil {
			p.err = fmt.Errorf("arg2 not an int: %s", tokens[2])
			return false
		}

		p.arg2 = arg2
		// default:
		// 	p.err = fmt.Errorf("unrecognized command type: %d", cmdType)
		// 	return false
	}

	return true
}

func stripCommentsAndWhitespace(text string) string {
	if i := strings.Index(text, commentPrefix); i >= 0 {
		text = text[:i]
	}
	return strings.TrimSpace(text)
}

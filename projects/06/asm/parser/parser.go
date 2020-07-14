package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	commentPrefix = "//"
)

var (
	cRegexp = regexp.MustCompile(`^((A?M?D?)\=)?([DAM]\+1|[DAM]\-1|D\+[AM]|D\-[AM]|[AM]\-D|[01DAM]|D\&[AM]|D\|[AM]|\-[01DAM]|\![DAM]);?(JGT|JEQ|JGE|JLT|JNE|JLE|JMP)?$`)
	aRegexp = regexp.MustCompile(`^\@([A-Za-z_.$:]+[A-Za-z0-9_.$:]*)|([0-9]+)$`)
	lRegexp = regexp.MustCompile(`^\(([A-Za-z_.$:]+[A-Za-z0-9_.$:]*)\)$`)

	compMap map[string]string = map[string]string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",
		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}
	jumpMap map[string]string = map[string]string{
		"":    "000",
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}
)

type Parser struct {
	rs      io.ReadSeeker
	scanner *bufio.Scanner
	err     error
	cmdType CommandType
	symbol  string
	decimal uint16
	dest    string
	comp    string
	jump    string
}

func New(rs io.ReadSeeker) *Parser {
	return &Parser{
		rs:      rs,
		scanner: bufio.NewScanner(rs),
	}
}

func (p *Parser) Rewind() {
	p.rs.Seek(0, 0)
	p.scanner = bufio.NewScanner(p.rs)
}

func (p *Parser) Err() error {
	return p.err
}

func (p *Parser) CommandType() CommandType {
	return p.cmdType
}

func (p *Parser) Symbol() string {
	return p.symbol
}

func (p *Parser) Decimal() uint16 {
	return p.decimal
}

func (p *Parser) Dest() string {
	return p.dest
}

func (p *Parser) Comp() string {
	return p.comp
}

func (p *Parser) Jump() string {
	return p.jump
}

func (p *Parser) resetProps() {
	p.cmdType = CommandTypeUnknown
	p.symbol = ""
	p.decimal = 0
	p.dest = ""
	p.comp = ""
	p.jump = ""
}

// Parses the next line
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

	switch {
	case cRegexp.Match([]byte(text)):
		p.cmdType = CommandTypeC

		matches := cRegexp.FindStringSubmatch(text)
		if len(matches) != 5 {
			// TODO: Return error
		}
		p.dest = parseDest(matches[2])
		p.comp = parseComp(matches[3])
		p.jump = parseJump(matches[4])
	case aRegexp.Match([]byte(text)):
		p.cmdType = CommandTypeA

		matches := aRegexp.FindStringSubmatch(text)
		if matches[1] != "" { // Symbol
			p.symbol = matches[1]
		}
		if matches[2] != "" { // Decimal
			n, _ := strconv.Atoi(matches[2])
			p.decimal = uint16(n)
		}
	case lRegexp.Match([]byte(text)):
		p.cmdType = CommandTypeL

		matches := lRegexp.FindStringSubmatch(text)
		p.symbol = matches[1]
	default:
		// TODO: Throw error. Unknown command.
	}

	return true
}

func stripCommentsAndWhitespace(text string) string {
	if i := strings.Index(text, commentPrefix); i >= 0 {
		text = text[:i]
	}
	return strings.TrimSpace(text)
}

// Assumes text already in valid format
func parseDest(text string) string {
	var code uint8
	for _, dest := range text {
		switch dest {
		case 'A':
			code += 4
		case 'D':
			code += 2
		case 'M':
			code += 1
		}
	}
	return fmt.Sprintf("%03b", code)
}

func parseComp(text string) string {
	return compMap[text]
}

func parseJump(text string) string {
	return jumpMap[text]
}

type CommandType uint8

const (
	CommandTypeUnknown CommandType = iota
	CommandTypeA
	CommandTypeC
	CommandTypeL
)

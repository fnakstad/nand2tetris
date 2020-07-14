package parser

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_parseCCommand(t *testing.T) {
	tests := map[string]string{
		"AMD=D|A;JMP": "",
		"MD=D-1":      "",
		"D&A;JMP":     "",
	}

	for in, exp := range tests {
		out := parseCCommand(in)
		assert.Equal(t, out, exp)
	}
}

func Test_parseDest(t *testing.T) {
	tests := map[string]string{
		"A":   "100",
		"M":   "010",
		"D":   "001",
		"AMD": "111",
		"AD":  "101",
		"MD":  "011",
		"":    "000",
	}

	for in, exp := range tests {
		out := parseDest(in)
		assert.Equal(t, out, exp)
	}
}

func Test_parseComp(t *testing.T) {
	tests := map[string]string{
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

	for in, exp := range tests {
		out := parseComp(in)
		assert.Equal(t, out, exp)
	}
}

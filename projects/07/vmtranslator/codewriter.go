package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

var (
	asmPush = []string{
		"// push",
		"@%d",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"",
	}
	asmArithmetic = map[string][]string{
		"add": []string{
			"// add",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=D+M",
			"",
		},
		"sub": []string{
			"// sub",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=M-D",
			"",
		},
		"neg": []string{
			"// neg",
			"@SP",
			"A=M-1",
			"M=-M",
			"",
		},
		"eq": []string{
			"// eq",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"D=M-D",
			"M=-1",
			"@EQ_%[1]d",
			"D;JEQ",
			"@SP",
			"A=M-1",
			"M=0",
			"(EQ_%[1]d)",
			"",
		},
		"gt": []string{
			"// gt",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"D=M-D",
			"M=-1",
			"@GT_%[1]d",
			"D;JGT",
			"@SP",
			"A=M-1",
			"M=0",
			"(GT_%[1]d)",
			"",
		},
		"lt": []string{
			"// lt",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"D=M-D",
			"M=-1",
			"@LT_%[1]d",
			"D;JLT",
			"@SP",
			"A=M-1",
			"M=0",
			"(LT_%[1]d)",
			"",
		},
		"and": []string{
			"// and",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=D&M",
			"",
		},
		"or": []string{
			"// or",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=D|M",
			"",
		},
		"not": []string{
			"// not",
			"@SP",
			"A=M-1",
			"M=!M",
			"",
		},
	}
)

type CodeWriter struct {
	w  io.Writer
	lc map[string]uint8
}

func NewCodeWriter(w io.Writer) *CodeWriter {
	return &CodeWriter{
		w:  w,
		lc: make(map[string]uint8),
	}
}

func (cw *CodeWriter) WriteArithmetic(cmd string) {
	var asm string
	switch cmd {
	case "gt", "lt", "eq":
		count := cw.lc[cmd]
		cw.lc[cmd]++
		asm = fmt.Sprintf(strings.Join(asmArithmetic[cmd], "\n"), count)
	default:
		// TODO: check for non-existent key
		asm = strings.Join(asmArithmetic[cmd], "\n")
	}

	_, err := io.WriteString(cw.w, asm)
	if err != nil {
		log.Fatalf("nooo: %v", err)
	}
}

func (cw *CodeWriter) WritePushPop(cmdType CommandType, segment string, index int) {
	switch cmdType {
	case CommandTypePush:
		if segment == "constant" {
			asm := fmt.Sprintf(strings.Join(asmPush, "\n"), index)
			_, err := io.WriteString(cw.w, asm)
			if err != nil {
				log.Fatalf("nooo: %v", err)
			}
		}
	case CommandTypePop:
		log.Println("pop")
	}
}

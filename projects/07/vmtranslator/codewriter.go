package main

import (
	"fmt"
	"io"
	"strings"
)

var (
	asmMap = map[CommandType][]string{
		CommandTypePush: []string{
			"// push",
			"@%d",
			"D=A",
			"@SP",
			"A=M",
			"M=D",
			"@SP",
			"M=M+1",
			"",
		},
		CommandTypeAdd: []string{
			"// add",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=D+M",
			"",
		},
		CommandTypeSub: []string{
			"// sub",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=M-D",
			"",
		},
		CommandTypeNeg: []string{
			"// neg",
			"@SP",
			"A=M-1",
			"M=-M",
			"",
		},
		CommandTypeEq: []string{
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
		CommandTypeGt: []string{
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
		CommandTypeLt: []string{
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
		CommandTypeAnd: []string{
			"// and",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=D&M",
			"",
		},
		CommandTypeOr: []string{
			"// or",
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=D|M",
			"",
		},
		CommandTypeNot: []string{
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
	lc map[CommandType]uint8
}

func NewCodeWriter(w io.Writer) *CodeWriter {
	return &CodeWriter{
		w:  w,
		lc: make(map[CommandType]uint8),
	}
}

func (cw *CodeWriter) WriteArithmetic(cmdType CommandType) error {
	var asm string
	switch cmdType {
	case CommandTypeGt, CommandTypeLt, CommandTypeEq:
		count := cw.lc[cmdType]
		cw.lc[cmdType]++
		asm = fmt.Sprintf(strings.Join(asmMap[cmdType], "\n"), count)
	default:
		// TODO: check for non-existent key
		asm = strings.Join(asmMap[cmdType], "\n")
	}

	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WritePushPop(cmdType CommandType, segmentType SegmentType, index int) error {
	asmStrings := asmMap[cmdType]
	switch segmentType {
	case SegmentTypeConstant:
		asm := fmt.Sprintf(strings.Join(asmStrings, "\n"), index)
		return cw.writeCommand(asm)
	}

	return nil
}

func (cw *CodeWriter) writeCommand(cmd string) error {
	_, err := io.WriteString(cw.w, cmd)
	return err
}

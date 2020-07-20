package main

import (
	"fmt"
	"io"
	"strings"
)

var (
	asmPushConstant = []string{
		"// push constant",
		"@%d",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"",
	}
	asmPushLATT = []string{
		"// push",
		"@%[2]s",
		"D=M",
		"@%[1]d",
		"A=A+D",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
	asmPushTP = []string{
		"// push",
		"@%[1]d",
		"D=A",
		"@%[2]s",
		"A=A+D",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
	asmPopLATT = []string{
		"// pop",
		"@%[2]s",
		"D=M",
		"@%[1]d",
		"D=A+D",
		"@R13",
		"M=D",
		"@SP",
		"AM=M-1",
		"D=M",
		"@R13",
		"A=M",
		"M=D",
		"",
	}
	asmPopTP = []string{
		"// pop",
		"@%[1]d",
		"D=A",
		"@%[2]s",
		"D=A+D",
		"@R13",
		"M=D",
		"@SP",
		"AM=M-1",
		"D=M",
		"@R13",
		"A=M",
		"M=D",
		"",
	}
	asmArithmeticMap = map[CommandType][]string{
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
		asm = fmt.Sprintf(strings.Join(asmArithmeticMap[cmdType], "\n"), count)
	default:
		// TODO: check for non-existent key
		asm = strings.Join(asmArithmeticMap[cmdType], "\n")
	}

	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WritePushPop(cmdType CommandType, segmentType SegmentType, index int) error {
	// log.Printf("%s %s %d", cmdType, segmentType, index)
	switch cmdType {
	case CommandTypePush:
		switch segmentType {
		case SegmentTypeConstant:
			asm := fmt.Sprintf(strings.Join(asmPushConstant, "\n"), index)
			return cw.writeCommand(asm)
		case SegmentTypeLocal, SegmentTypeArgument, SegmentTypeThis, SegmentTypeThat:
			segmentBase, ok := segmentBaseMap[segmentType]
			if !ok {
				return fmt.Errorf("can't find segment base for type %v", segmentBase)
			}
			asm := fmt.Sprintf(strings.Join(asmPushLATT, "\n"), index, segmentBase)
			return cw.writeCommand(asm)
		case SegmentTypeTemp, SegmentTypePointer:
			segmentBase, ok := segmentBaseMap[segmentType]
			if !ok {
				return fmt.Errorf("can't find segment base for type %v", segmentBase)
			}
			asm := fmt.Sprintf(strings.Join(asmPushTP, "\n"), index, segmentBase)
			return cw.writeCommand(asm)
		}
	case CommandTypePop:
		switch segmentType {
		case SegmentTypeLocal, SegmentTypeArgument, SegmentTypeThis, SegmentTypeThat:
			segmentBase, ok := segmentBaseMap[segmentType]
			if !ok {
				return fmt.Errorf("can't find segment base for type %v", segmentBase)
			}
			asm := fmt.Sprintf(strings.Join(asmPopLATT, "\n"), index, segmentBase)
			return cw.writeCommand(asm)
		case SegmentTypeTemp, SegmentTypePointer:
			segmentBase, ok := segmentBaseMap[segmentType]
			if !ok {
				return fmt.Errorf("can't find segment base for type %v", segmentBase)
			}
			asm := fmt.Sprintf(strings.Join(asmPopTP, "\n"), index, segmentBase)
			return cw.writeCommand(asm)
		}
	}

	return nil
}

func (cw *CodeWriter) writeCommand(cmd string) error {
	_, err := io.WriteString(cw.w, cmd)
	return err
}

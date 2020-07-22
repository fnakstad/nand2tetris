package main

import "strings"

var (
	// Arithmetic
	asmAdd = []string{
		"// add",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D+M",
		"",
	}
	asmSub = []string{
		"// sub",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=M-D",
		"",
	}
	asmNeg = []string{
		"// neg",
		"@SP",
		"A=M-1",
		"M=-M",
		"",
	}
	asmEq = []string{
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
	}
	asmGt = []string{
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
	}
	asmLt = []string{
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
	}
	asmAnd = []string{
		"// and",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D&M",
		"",
	}
	asmOr = []string{
		"// or",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D|M",
		"",
	}
	asmNot = []string{
		"// not",
		"@SP",
		"A=M-1",
		"M=!M",
		"",
	}

	// Stack/memory access
	asmPushDToStack = []string{
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"",
	}
	asmPushStatic = []string{
		"// push static",
		"@%[1]s",
		"D=M",
		strings.Join(asmPushDToStack, "\n"),
	}
	asmPushConstant = []string{
		"// push constant",
		"@%d",
		"D=A",
		strings.Join(asmPushDToStack, "\n"),
	}
	asmPushLATT = []string{
		"// push",
		"@%[2]s",
		"D=M",
		"@%[1]d",
		"A=A+D",
		"D=M",
		strings.Join(asmPushDToStack, "\n"),
	}
	asmPushTP = []string{
		"// push",
		"@%[1]d",
		"D=A",
		"@%[2]s",
		"A=A+D",
		"D=M",
		strings.Join(asmPushDToStack, "\n"),
	}
	asmPopStatic = []string{
		"// pop static",
		"@SP",
		"AM=M-1",
		"D=M",
		"@%[1]s",
		"M=D",
		"",
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

	// Control flow
	asmLabel = []string{
		"// label",
		"(%[1]s)",
		"",
	}
	asmIf = []string{
		"// if",
		"@SP",
		"AM=M-1",
		"D=M",
		"@%[1]s",
		"D;JGT",
		"",
	}
	asmGoto = []string{
		"// goto",
		"@%[1]s",
		"0;JMP",
	}

	// Functions
	asmFunction = []string{
		"// function",
		"(%[1]s)",
		"%[2]s",
	}
	asmReturn = []string{
		"// return",
		"",
		"",
	}
	asmCall = []string{
		"// call",
		"@%[1]s", // push return address
		"D=A",
		strings.Join(asmPushDToStack, "\n"),
		"",
	}
)

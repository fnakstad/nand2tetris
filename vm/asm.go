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
		"@LCL", // current frame
		"D=M",
		"@R13",
		"M=D",
		"@5", // return value
		"A=D-A",
		"D=M",
		"@R14",
		"M=D",
		"@SP", // reposition return value
		"A=M-1",
		"D=M",
		"@ARG",
		"A=M",
		"M=D",
		"D=A+1", // reposition SP
		"@SP",
		"M=D",
		"@R13", // reposition THAT
		"AM=M-1",
		"D=M",
		"@THAT",
		"M=D",
		"@R13", // reposition THIS
		"AM=M-1",
		"D=M",
		"@THIS",
		"M=D",
		"@R13", // reposition ARG
		"AM=M-1",
		"D=M",
		"@ARG",
		"M=D",
		"@R13", // reposition LCL
		"AM=M-1",
		"D=M",
		"@LCL",
		"M=D",
		"@R14", // go to return address
		"A=M",
		"0;JMP",
		"",
	}
	asmCall = []string{
		"// call",
		"@%[1]s", // push return address
		"D=A",
		strings.Join(asmPushDToStack, "\n"),
		"@LCL", // push LCL
		"D=A",
		strings.Join(asmPushDToStack, "\n"),
		"@ARG", // push ARG
		"D=A",
		strings.Join(asmPushDToStack, "\n"),
		"@THIS", // push THIS
		"D=A",
		strings.Join(asmPushDToStack, "\n"),
		"@THAT", // push THAT
		"D=A",
		strings.Join(asmPushDToStack, "\n"),
		"@SP", // reposition ARG (SP-n-5)
		"D=M",
		"@%[3]",
		"D=D-A",
		"@5",
		"D=D-A",
		"@ARG",
		"M=D",
		"@SP", // reposition LCL
		"D=M",
		"@LCL",
		"M=D",
		"@%[2]s", // jump to function
		"0;JMP",
		"(@%[1]s)", // label for return address
		"",
	}
)

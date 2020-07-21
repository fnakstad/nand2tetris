package main

var (
	asmPushStatic = []string{
		"// push static",
		"@%[1]s",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"",
	}
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

	// Arithemtic
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
)

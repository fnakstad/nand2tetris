// push
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@EQ_0
D;JEQ
@SP
A=M-1
M=0
(EQ_0)
// push
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@EQ_1
D;JEQ
@SP
A=M-1
M=0
(EQ_1)
// push
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@EQ_2
D;JEQ
@SP
A=M-1
M=0
(EQ_2)
// push
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@LT_0
D;JLT
@SP
A=M-1
M=0
(LT_0)
// push
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@LT_1
D;JLT
@SP
A=M-1
M=0
(LT_1)
// push
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@LT_2
D;JLT
@SP
A=M-1
M=0
(LT_2)
// push
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@GT_0
D;JGT
@SP
A=M-1
M=0
(GT_0)
// push
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@GT_1
D;JGT
@SP
A=M-1
M=0
(GT_1)
// push
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@GT_2
D;JGT
@SP
A=M-1
M=0
(GT_2)
// push
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@53
D=A
@SP
A=M
M=D
@SP
M=M+1
// add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// push
@112
D=A
@SP
A=M
M=D
@SP
M=M+1
// sub
@SP
AM=M-1
D=M
A=A-1
M=M-D
// neg
@SP
A=M-1
M=-M
// and
@SP
AM=M-1
D=M
A=A-1
M=D&M
// push
@82
D=A
@SP
A=M
M=D
@SP
M=M+1
// or
@SP
AM=M-1
D=M
A=A-1
M=D|M
// not
@SP
A=M-1
M=!M

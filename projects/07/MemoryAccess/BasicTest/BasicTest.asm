// push constant
@10
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@LCL
D=M
@0
D=A+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// push constant
@21
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant
@22
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@ARG
D=M
@2
D=A+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// pop
@ARG
D=M
@1
D=A+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// push constant
@36
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THIS
D=M
@6
D=A+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// push constant
@42
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant
@45
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THAT
D=M
@5
D=A+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// pop
@THAT
D=M
@2
D=A+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// push constant
@510
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@6
D=A
@R5
D=A+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// push
@LCL
D=M
@0
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1// push
@THAT
D=M
@5
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1// add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// push
@ARG
D=M
@1
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1// sub
@SP
AM=M-1
D=M
A=A-1
M=M-D
// push
@THIS
D=M
@6
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1// push
@THIS
D=M
@6
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1// add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// sub
@SP
AM=M-1
D=M
A=A-1
M=M-D
// push
@6
D=A
@R5
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1// add
@SP
AM=M-1
D=M
A=A-1
M=D+M

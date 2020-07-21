// push constant
@3030
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@0
D=A
@THIS
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
@3040
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@1
D=A
@THIS
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
@32
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THIS
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
@46
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THAT
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
// push
@0
D=A
@THIS
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1// push
@1
D=A
@THIS
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
@THIS
D=M
@2
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
@THAT
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

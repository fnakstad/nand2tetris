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
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THAT
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
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THAT
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
// push
@ARG
D=M
@0
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push constant
@2
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
// pop
@ARG
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
// label
(global$MAIN_LOOP_START)
// push
@ARG
D=M
@0
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// if
@SP
AM=M-1
D=M
@global$COMPUTE_ELEMENT
D;JGT
// goto
@global$END_PROGRAM
0;JMP// label
(global$COMPUTE_ELEMENT)
// push
@THAT
D=M
@0
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push
@THAT
D=M
@1
A=A+D
D=M
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
// push
@1
D=A
@THIS
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push constant
@1
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
// push
@ARG
D=M
@0
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push constant
@1
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
// pop
@ARG
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
// goto
@global$MAIN_LOOP_START
0;JMP// label
(global$END_PROGRAM)

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
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
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
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
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
@2
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=M-D
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
(global$MAIN_LOOP_START)
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
@SP
AM=M-1
D=M
@global$COMPUTE_ELEMENT
D;JNE
@global$END_PROGRAM
0;JMP
(global$COMPUTE_ELEMENT)
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
@SP
AM=M-1
D=M
A=A-1
M=D+M
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
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=D+M
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
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=M-D
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
@global$MAIN_LOOP_START
0;JMP
(global$END_PROGRAM)

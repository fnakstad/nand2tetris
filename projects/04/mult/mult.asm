// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

// Initialize
@R2
M=0
@R1
D=M
@i
M=D

(LOOP)
    // Loop check
    @i
    D=M
    @END
    D;JLE

    // Decrement R1
    @i
    D=M
    @1
    D=D-A
    @i
    M=D

    // Add R0 to R2
    @R2
    D=M
    @R0
    A=M
    D=D+A
    @R2
    M=D

    @LOOP
    0;JMP
(END)
    @END
    0;JMP

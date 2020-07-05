// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

// Continuously loop through screen
(INIT)
    // Initialize i
    @SCREEN
    D=A
    @i
    M=D

    // Initialize max
    @8192 // 32 * 256
    D=A+D
    @max
    M=D
(LOOP)
    // TODO: Set D to color based on keyboard input
    @KBD
    D=M
    @SETBLACK
    D;JGT

(DRAW)
    // Set current position to color stored in D
    @i
    A=M
    M=D

    // Increment i
    D=A+1
    @i
    M=D

    // If end reached, start over from start
    @max
    D=D-M
    @INIT
    D;JGT

    // Otherwise, keep looping
    @LOOP
    0;JMP

(SETBLACK)
    D=0
    D=!D
    @DRAW
    0;JMP
CPU_EMU=../tools/CPUEmulator.sh
GO_RUN=go run .
PROJECTS_DIR=../projects

.PHONY: all
all: translate test

.PHONY: translate
translate:
	${GO_RUN} ${PROJECTS_DIR}/07/StackArithmetic/SimpleAdd/SimpleAdd.vm ${PROJECTS_DIR}/07/StackArithmetic/SimpleAdd/SimpleAdd.asm
	${GO_RUN} ${PROJECTS_DIR}/07/StackArithmetic/StackTest/StackTest.vm ${PROJECTS_DIR}/07/StackArithmetic/StackTest/StackTest.asm
	${GO_RUN} ${PROJECTS_DIR}/07/MemoryAccess/BasicTest/BasicTest.vm ${PROJECTS_DIR}/07/MemoryAccess/BasicTest/BasicTest.asm
	${GO_RUN} ${PROJECTS_DIR}/07/MemoryAccess/PointerTest/PointerTest.vm ${PROJECTS_DIR}/07/MemoryAccess/PointerTest/PointerTest.asm
	${GO_RUN} ${PROJECTS_DIR}/07/MemoryAccess/StaticTest/StaticTest.vm ${PROJECTS_DIR}/07/MemoryAccess/StaticTest/StaticTest.asm
	${GO_RUN} ${PROJECTS_DIR}/08/ProgramFlow/BasicLoop/BasicLoop.vm ${PROJECTS_DIR}/08/ProgramFlow/BasicLoop/BasicLoop.asm
	${GO_RUN} ${PROJECTS_DIR}/08/ProgramFlow/FibonacciSeries/FibonacciSeries.vm ${PROJECTS_DIR}/08/ProgramFlow/FibonacciSeries/FibonacciSeries.asm
	${GO_RUN} ${PROJECTS_DIR}/08/FunctionCalls/SimpleFunction/SimpleFunction.vm ${PROJECTS_DIR}/08/FunctionCalls/SimpleFunction/SimpleFunction.asm

.PHONY: test
test:
	${CPU_EMU} ${PROJECTS_DIR}/07/StackArithmetic/SimpleAdd/SimpleAdd.tst
	${CPU_EMU} ${PROJECTS_DIR}/07/StackArithmetic/StackTest/StackTest.tst
	${CPU_EMU} ${PROJECTS_DIR}/07/MemoryAccess/BasicTest/BasicTest.tst
	${CPU_EMU} ${PROJECTS_DIR}/07/MemoryAccess/PointerTest/PointerTest.tst
	${CPU_EMU} ${PROJECTS_DIR}/07/MemoryAccess/StaticTest/StaticTest.tst
	${CPU_EMU} ${PROJECTS_DIR}/08/ProgramFlow/BasicLoop/BasicLoop.tst
	${CPU_EMU} ${PROJECTS_DIR}/08/ProgramFlow/FibonacciSeries/FibonacciSeries.tst
	${CPU_EMU} ${PROJECTS_DIR}/08/FunctionCalls/SimpleFunction/SimpleFunction.tst
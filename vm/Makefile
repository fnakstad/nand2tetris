CPU_EMU=../tools/CPUEmulator.sh
GO_RUN=go run .
PROJECTS_DIR=../projects

.PHONY: all
all: translate test

.PHONY: translate
translate:
	${GO_RUN} -in=${PROJECTS_DIR}/07/StackArithmetic/SimpleAdd -out=${PROJECTS_DIR}/07/StackArithmetic/SimpleAdd/SimpleAdd.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/07/StackArithmetic/StackTest -out=${PROJECTS_DIR}/07/StackArithmetic/StackTest/StackTest.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/07/MemoryAccess/BasicTest -out=${PROJECTS_DIR}/07/MemoryAccess/BasicTest/BasicTest.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/07/MemoryAccess/PointerTest -out=${PROJECTS_DIR}/07/MemoryAccess/PointerTest/PointerTest.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/07/MemoryAccess/StaticTest -out=${PROJECTS_DIR}/07/MemoryAccess/StaticTest/StaticTest.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/08/ProgramFlow/BasicLoop -out=${PROJECTS_DIR}/08/ProgramFlow/BasicLoop/BasicLoop.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/08/ProgramFlow/FibonacciSeries -out=${PROJECTS_DIR}/08/ProgramFlow/FibonacciSeries/FibonacciSeries.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/08/FunctionCalls/SimpleFunction -out=${PROJECTS_DIR}/08/FunctionCalls/SimpleFunction/SimpleFunction.asm -bootstrap=false
	${GO_RUN} -in=${PROJECTS_DIR}/08/FunctionCalls/NestedCall -out=${PROJECTS_DIR}/08/FunctionCalls/NestedCall/NestedCall.asm
	${GO_RUN} -in=${PROJECTS_DIR}/08/FunctionCalls/FibonacciElement -out=${PROJECTS_DIR}/08/FunctionCalls/FibonacciElement/FibonacciElement.asm
	${GO_RUN} -in=${PROJECTS_DIR}/08/FunctionCalls/StaticsTest -out=${PROJECTS_DIR}/08/FunctionCalls/StaticsTest/StaticsTest.asm

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
	${CPU_EMU} ${PROJECTS_DIR}/08/FunctionCalls/NestedCall/NestedCall.tst
	${CPU_EMU} ${PROJECTS_DIR}/08/FunctionCalls/FibonacciElement/FibonacciElement.tst
	${CPU_EMU} ${PROJECTS_DIR}/08/FunctionCalls/StaticsTest/StaticsTest.tst

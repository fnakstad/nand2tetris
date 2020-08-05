CMP_TOOL=../tools/TextComparer.sh
GO_RUN=go run .
PROJECTS_DIR=../projects

.PHONY: all
all: compile test

.PHONY: compile
compile:
	${GO_RUN} -in ${PROJECTS_DIR}/10/ArrayTest
	${GO_RUN} -in ${PROJECTS_DIR}/10/ExpressionLessSquare
	${GO_RUN} -in ${PROJECTS_DIR}/10/Square

.PHONY: test
test:
	${CMP_TOOL} ${PROJECTS_DIR}/10/ArrayTest/MainT_.xml ${PROJECTS_DIR}/10/ArrayTest/MainT.xml
	${CMP_TOOL} ${PROJECTS_DIR}/10/Square/MainT_.xml ${PROJECTS_DIR}/10/Square/MainT.xml
	${CMP_TOOL} ${PROJECTS_DIR}/10/Square/SquareT_.xml ${PROJECTS_DIR}/10/Square/SquareT.xml
	${CMP_TOOL} ${PROJECTS_DIR}/10/Square/SquareGameT_.xml ${PROJECTS_DIR}/10/Square/SquareGameT.xml
	${CMP_TOOL} ${PROJECTS_DIR}/10/ExpressionLessSquare/MainT_.xml ${PROJECTS_DIR}/10/ExpressionLessSquare/MainT.xml
	${CMP_TOOL} ${PROJECTS_DIR}/10/ExpressionLessSquare/SquareT_.xml ${PROJECTS_DIR}/10/ExpressionLessSquare/SquareT.xml
	${CMP_TOOL} ${PROJECTS_DIR}/10/ExpressionLessSquare/SquareGameT_.xml ${PROJECTS_DIR}/10/ExpressionLessSquare/SquareGameT.xml
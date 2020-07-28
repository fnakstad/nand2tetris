package main

import (
	"fmt"
	"io"
	"strings"
)

// TODO: namespace user provided labels to prevent collisions
// const (
// 	userLabelPrefix = "USR_"
// )

var (
	asmPushPopMap = map[CommandType]map[SegmentType][]string{
		CommandTypePush: map[SegmentType][]string{
			SegmentTypeStatic:   asmPushStatic,
			SegmentTypeConstant: asmPushConstant,
			SegmentTypeLocal:    asmPushLATT,
			SegmentTypeArgument: asmPushLATT,
			SegmentTypeThis:     asmPushLATT,
			SegmentTypeThat:     asmPushLATT,
			SegmentTypeTemp:     asmPushTP,
			SegmentTypePointer:  asmPushTP,
		},
		CommandTypePop: map[SegmentType][]string{
			SegmentTypeStatic:   asmPopStatic,
			SegmentTypeLocal:    asmPopLATT,
			SegmentTypeArgument: asmPopLATT,
			SegmentTypeThis:     asmPopLATT,
			SegmentTypeThat:     asmPopLATT,
			SegmentTypeTemp:     asmPopTP,
			SegmentTypePointer:  asmPopTP,
		},
	}
	asmArithmeticMap = map[CommandType][]string{
		CommandTypeAdd: asmAdd,
		CommandTypeSub: asmSub,
		CommandTypeNeg: asmNeg,
		CommandTypeEq:  asmEq,
		CommandTypeGt:  asmGt,
		CommandTypeLt:  asmLt,
		CommandTypeAnd: asmAnd,
		CommandTypeOr:  asmOr,
		CommandTypeNot: asmNot,
	}

	callNumMap = map[string]int{}
)

type CodeWriter struct {
	w  io.Writer
	lc map[CommandType]uint8
}

func NewCodeWriter(w io.Writer) *CodeWriter {
	return &CodeWriter{
		w:  w,
		lc: make(map[CommandType]uint8),
	}
}

func (cw *CodeWriter) WriteBootstrap() error {
	asm := strings.Join(asmBootstrap, "\n")
	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WriteArithmetic(cmdType CommandType) error {
	var asm string
	asmStrings, ok := asmArithmeticMap[cmdType]
	if !ok {
		return fmt.Errorf("no asm handler found for %v", cmdType)
	}
	switch cmdType {
	case CommandTypeGt, CommandTypeLt, CommandTypeEq:
		count := cw.lc[cmdType]
		cw.lc[cmdType]++
		asm = fmt.Sprintf(strings.Join(asmStrings, "\n"), count)
	default:
		asm = strings.Join(asmStrings, "\n")
	}

	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WritePushPop(cmdType CommandType, segmentType SegmentType, index int, fid string) error {
	var asm string
	asmStrings, ok := asmPushPopMap[cmdType][segmentType]
	if !ok {
		return fmt.Errorf("no asm handler found for %v", cmdType)
	}

	switch segmentType {
	case SegmentTypeConstant:
		asm = fmt.Sprintf(strings.Join(asmStrings, "\n"), index)
	case SegmentTypeStatic:
		varName := fmt.Sprintf("%s.%d", fid, index)
		asm = fmt.Sprintf(strings.Join(asmStrings, "\n"), varName)
	default:
		segmentBase, ok := segmentBaseMap[segmentType]
		if !ok {
			return fmt.Errorf("can't find segment base for type %v", segmentBase)
		}
		asm = fmt.Sprintf(strings.Join(asmStrings, "\n"), index, segmentBase)
	}

	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WriteLabel(funcName, label string) error {
	nsLabel := getNamespacedLabel(funcName, label)
	asm := fmt.Sprintf(strings.Join(asmLabel, "\n"), nsLabel)
	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WriteIf(funcName, label string) error {
	nsLabel := getNamespacedLabel(funcName, label)
	asm := fmt.Sprintf(strings.Join(asmIf, "\n"), nsLabel)
	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WriteGoto(funcName, label string) error {
	nsLabel := getNamespacedLabel(funcName, label)
	asm := fmt.Sprintf(strings.Join(asmGoto, "\n"), nsLabel)
	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WriteFunction(funcName string, numLocals int) error {
	lclinit := make([]string, numLocals)
	for i := 0; i < numLocals; i++ {
		lclinit[i] = fmt.Sprintf(strings.Join(asmPushLATT, "\n"), i, "LCL")
	}

	asm := fmt.Sprintf(strings.Join(asmFunction, "\n"), funcName, strings.Join(lclinit, ""))
	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WriteReturn() error {
	asm := strings.Join(asmReturn, "\n")
	return cw.writeCommand(asm)
}

func (cw *CodeWriter) WriteCall(funcName string, numArgs int) error {
	// Get callNum from map
	callNum := callNumMap[funcName]
	callNumMap[funcName]++
	returnAddress := fmt.Sprintf("return_%s_%d", funcName, callNum)
	asm := fmt.Sprintf(strings.Join(asmCall, "\n"), returnAddress, funcName, numArgs)
	return cw.writeCommand(asm)
}

func getNamespacedLabel(funcName, label string) string {
	return fmt.Sprintf("%s$%s", funcName, label)
}

func (cw *CodeWriter) writeCommand(cmd string) error {
	_, err := io.WriteString(cw.w, cmd)
	return err
}

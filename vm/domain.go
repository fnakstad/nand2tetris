package main

type CommandType string

const (
	CommandTypeUnknown CommandType = ""

	// Arithmetic
	CommandTypeAdd = "add"
	CommandTypeSub = "sub"
	CommandTypeNeg = "neg"
	CommandTypeEq  = "eq"
	CommandTypeGt  = "gt"
	CommandTypeLt  = "lt"
	CommandTypeAnd = "and"
	CommandTypeOr  = "or"
	CommandTypeNot = "not"

	// Others
	CommandTypePush     = "push"
	CommandTypePop      = "pop"
	CommandTypeLabel    = "label"
	CommandTypeGoto     = "goto"
	CommandTypeIf       = "if-goto"
	CommandTypeFunction = "function"
	CommandTypeReturn   = "return"
	CommandTypeCall     = "call"
)

type SegmentType string

const (
	SegmentTypeUnknown  SegmentType = ""
	SegmentTypeArgument             = "argument"
	SegmentTypeLocal                = "local"
	SegmentTypeStatic               = "static"
	SegmentTypeConstant             = "constant"
	SegmentTypeThis                 = "this"
	SegmentTypeThat                 = "that"
	SegmentTypePointer              = "pointer"
	SegmentTypeTemp                 = "temp"
)

var (
	arithmeticCommand = map[CommandType]struct{}{
		CommandTypeAdd: struct{}{},
		CommandTypeSub: struct{}{},
		CommandTypeNeg: struct{}{},
		CommandTypeEq:  struct{}{},
		CommandTypeGt:  struct{}{},
		CommandTypeLt:  struct{}{},
		CommandTypeAnd: struct{}{},
		CommandTypeOr:  struct{}{},
		CommandTypeNot: struct{}{},
	}
	tokenCommandMap = map[string]CommandType{
		"add":      CommandTypeAdd,
		"sub":      CommandTypeSub,
		"neg":      CommandTypeNeg,
		"eq":       CommandTypeEq,
		"gt":       CommandTypeGt,
		"lt":       CommandTypeLt,
		"and":      CommandTypeAnd,
		"or":       CommandTypeOr,
		"not":      CommandTypeNot,
		"push":     CommandTypePush,
		"pop":      CommandTypePop,
		"label":    CommandTypeLabel,
		"goto":     CommandTypeGoto,
		"if-goto":  CommandTypeIf,
		"function": CommandTypeFunction,
		"return":   CommandTypeReturn,
		"call":     CommandTypeCall,
	}
	segmentBaseMap = map[SegmentType]string{
		SegmentTypeLocal:    "LCL", // pointer to segment
		SegmentTypeArgument: "ARG",
		SegmentTypeThis:     "THIS",
		SegmentTypeThat:     "THAT",
		SegmentTypeTemp:     "R5", // direct storage address
		SegmentTypePointer:  "THIS",
	}
)

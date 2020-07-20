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
	CommandTypeIf       = "if"
	CommandTypeFunction = "func"
	CommandTypeReturn   = "return"
	CommandTypeCall     = "call"
)

type SegmentType string

const (
	SegmentTypeUnknown  SegmentType = ""
	SegmentTypeConstant             = "constant"
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
	tokenCommandMapping = map[string]CommandType{
		"add":  CommandTypeAdd,
		"sub":  CommandTypeSub,
		"neg":  CommandTypeNeg,
		"eq":   CommandTypeEq,
		"gt":   CommandTypeGt,
		"lt":   CommandTypeLt,
		"and":  CommandTypeAnd,
		"or":   CommandTypeOr,
		"not":  CommandTypeNot,
		"push": CommandTypePush,
		"pop":  CommandTypePop,
	}
)

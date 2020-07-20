package main

type CommandType uint8

const (
	CommandTypeUnknown CommandType = iota
	CommandTypeArithmetic
	CommandTypePush
	CommandTypePop
	CommandTypeLabel
	CommandTypeGoto
	CommandTypeIf
	CommandTypeFunction
	CommandTypeReturn
	CommandTypeCall
)

var (
	tokenCommandMapping = map[string]CommandType{
		"add":  CommandTypeArithmetic,
		"sub":  CommandTypeArithmetic,
		"neg":  CommandTypeArithmetic,
		"eq":   CommandTypeArithmetic,
		"gt":   CommandTypeArithmetic,
		"lt":   CommandTypeArithmetic,
		"and":  CommandTypeArithmetic,
		"or":   CommandTypeArithmetic,
		"not":  CommandTypeArithmetic,
		"push": CommandTypePush,
		"pop":  CommandTypePop,
	}
)

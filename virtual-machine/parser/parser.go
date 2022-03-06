package parser

import "strings"

type CommandType int8
const (
	Arithmetic CommandType = 0
	Memory CommandType = 1
	Branching CommandType = 2
	FunctionStuff CommandType = 3

)
var arithmeticCommands = map[string]struct{} {
	"add": struct{}{},
	"sub": struct{}{},
	"neg": struct{}{},
	"eq": struct{}{},
	"gt": struct{}{},
	"lt": struct{}{},
	"and": struct{}{},
	"or": struct{}{},
	"not": struct{}{},
}

// Parses instruction into it's type and operands
func ParseLine(line string) (CommandType, string, string, string) {
	line = strings.Trim(line, " ")
	commandType := getCommandType(line)
	if commandType == Arithmetic {
		return Arithmetic, line, "", ""
	} else if commandType == Memory {
		split := strings.Split(line, " ")
		return Memory, split[0], split[1], split[2]
	}

	return -1, "", "", ""
}

// Return commands type. Types are Arithmetic, Memory, Branching or Function
func getCommandType(line string) CommandType {
	_, isArithmetic := arithmeticCommands[line]
	if isArithmetic {
		return Arithmetic
	} else if strings.Contains(line, "push") || strings.Contains(line, "pop") {
		return Memory
	}

	return -1
}

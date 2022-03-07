package code

import (
	"log"
	"strings"
)

// This is still not good
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
func (context *CodeContext) parseLine(line string) {
	_, isArithmetic := arithmeticCommands[line]
	if isArithmetic {
		context.currentTranslator = context.arithmeticTranslator
		context.currentCommand = line
	} else if strings.Contains(line, "push") || strings.Contains(line, "pop") {
		split := strings.Split(line, " ")
		context.currentTranslator = context.memorySegmentTranslator
		context.currentCommand = split[0]
		context.currentFirstArg = split[1]
		context.currentSecondArg = split[2]
	} else {
		log.Fatalln("This should be handled")
	}
}

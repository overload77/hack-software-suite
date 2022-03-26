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

// Parses instruction into it's type and operands and set's up context
func (context *CodeContext) parseLine(line string) {
	_, isArithmetic := arithmeticCommands[line]
	if isArithmetic {
		context.setTranslatorAndArgs(context.arithmeticTranslator, line)
	} else if strings.Contains(line, "push") || strings.Contains(line, "pop") {
		split := strings.Split(line, " ")
		context.setTranslatorAndArgs(context.memorySegmentTranslator, split[0], split[1], split[2])
	} else if isBranchingCommand(line) {
		split := strings.Split(line, " ")
		context.setTranslatorAndArgs(context.branchTranslator, split[0], split[1], "")
	} else if isFunctionCommand(line) {
		if split := strings.Split(line, " "); len(split) > 1 {
			context.setTranslatorAndArgs(context.functionTranslator, split[0], split[1], split[2])
		} else {
			context.setTranslatorAndArgs(context.functionTranslator, split[0])
		}
	} else {
		log.Fatalln("Unintegrated translator")
	}
}

// Sets the code context's current translator, command and optionally first and second args
func (context *CodeContext) setTranslatorAndArgs(translator Translator, command string,
												 args ...string) {
	context.currentTranslator = translator
	context.currentCommand = command
	if len(args) != 0 {
		context.currentFirstArg = args[0]
		context.currentSecondArg = args[1]
	}
}

func isBranchingCommand(line string) bool {
	branchingCommands := []string{"label", "goto", "if-goto"}
	for _, command := range branchingCommands {
		if strings.Contains(line, command) {
			return true
		}
	}
	return false
}

func isFunctionCommand(line string) bool {
	functionCommands := []string{"call", "function", "return"}
	for _, command := range functionCommands {
		if strings.Contains(line, command) {
			return true
		}
	}
	return false
}
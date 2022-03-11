package code

import (
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code/arithmetic"
	"github.com/overload77/hack-software-suite/virtual-machine/code/branch"
	"github.com/overload77/hack-software-suite/virtual-machine/code/memory"
)

type CodeContext struct {
	arithmeticTranslator *arithmetic.ArithmeticTranslator
	memorySegmentTranslator *memory.MemorySegmentTranslator
	branchTranslator *branch.BranchTranslator
	builder *strings.Builder
	vmFileName string
	currentTranslator Translator
	currentCommand string
	currentFirstArg string
	currentSecondArg string
}

type Translator interface {
	Translate(string, string, string)
}

func GetCodeContext(vmFileName string) *CodeContext {
	builder := &strings.Builder{}
	return &CodeContext {
		arithmeticTranslator: arithmetic.GetArithmeticTranslator(builder, vmFileName),
		memorySegmentTranslator: memory.GetMemorySegmentTranslator(builder, vmFileName),
		branchTranslator: branch.GetBranchTranslator(builder, vmFileName),
		builder: builder,
		vmFileName: vmFileName,
	}
}

func (context *CodeContext) TranslateCommand(command string) {
	if command, isValidCommand := trimLine(command); isValidCommand {
		context.parseLine(command)
		context.currentTranslator.Translate(
			context.currentCommand, context.currentFirstArg, context.currentSecondArg)
	}
}

func (context *CodeContext) GetCodeString() string {
	return context.builder.String()
}

func trimLine(line string) (string, bool) {
	line = strings.Trim(line, " ")
	if strings.HasPrefix(line, "//") || len(line) == 0 {
		return "", false
	} else if commentStart := strings.Index(line, "//"); commentStart != -1 {
		return strings.Trim(line[:commentStart], " "), true
	}

	return line, true
}
package code

import (
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code/arithmetic"
	"github.com/overload77/hack-software-suite/virtual-machine/code/memory"
)

type CodeContext struct {
	arithmeticTranslator *arithmetic.ArithmeticTranslator
	memorySegmentTranslator *memory.MemorySegmentTranslator
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
		builder: builder,
		vmFileName: vmFileName,
	}
}

func (context *CodeContext) TranslateCommand(command string) {
	context.parseLine(command)
	context.currentTranslator.Translate(
		context.currentCommand, context.currentFirstArg, context.currentSecondArg)
}

func (context *CodeContext) GetCodeString() string {
	return context.builder.String()
}

package code

import (
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code/memory"
	"github.com/overload77/hack-software-suite/virtual-machine/parser"
)

type CodeContext struct {
	arithmeticTranslator *ArithmeticTranslator
	memorySegmentTranslator *memory.MemorySegmentTranslator
	builder *strings.Builder
	vmFileName string
}

func GetCodeContext(vmFileName string) *CodeContext {
	builder := &strings.Builder{}
	return &CodeContext {
		arithmeticTranslator: GetArithmeticTranslator(builder, vmFileName),
		memorySegmentTranslator: memory.GetMemorySegmentTranslator(builder, vmFileName),
		builder: builder,
		vmFileName: vmFileName,
	}
}

func (context *CodeContext) TranslateCommand(commandType parser.CommandType,
		commandName string, firstCommandArg string, secondCommandArg string) {
	if commandType == parser.Arithmetic {
		context.translateArithmetic(commandName)
	} else if commandType == parser.Memory {
		context.translateMemory(commandName, firstCommandArg, secondCommandArg)
	}
}

func (context *CodeContext) GetCodeString() string {
	return context.builder.String()
}

func (context *CodeContext) translateArithmetic(commandName string) {
	context.arithmeticTranslator.Translate(commandName)
}

func (context *CodeContext) translateMemory(pushOrPop string, segment string, index string) {
	context.memorySegmentTranslator.Translate(pushOrPop, segment, index)
}

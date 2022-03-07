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
	currentTranslator interface{} // Create an intr that all translaters implement Translate()
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

// TODO: Hmm, let's add currentTranslator, commandName, firstArg and secondArg into context
// And let parser set those. That'll eliminate multiple switches. Then code context can just
// call
func (context *CodeContext) TranslateCommand(command string) {
	commandType, commandName, firstArg, secondArg := parser.ParseLine(command)
	switch commandType {
	case parser.Arithmetic:
		context.translateArithmetic(commandName)
	case parser.Memory:
		context.translateMemory(commandName, firstArg, secondArg)
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

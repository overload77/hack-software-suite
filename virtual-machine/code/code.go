package code

import (
	"log"
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code/memory"
)

type CodeContext struct {
	arithmeticTranslator *ArithmeticTranslator
	memorySegmentTranslator *memory.MemorySegmentTranslator
	builder *strings.Builder
	vmFileName string
}

func GetCodeContext(params ...interface{}) *CodeContext {
	startingBranchNum := getStartingBranchNumber(params...)
	vmFileName := params[0].(string)
	builder := &strings.Builder{}
	return &CodeContext {
		arithmeticTranslator: GetArithmeticTranslator(startingBranchNum, builder),
		memorySegmentTranslator: memory.GetMemorySegmentTranslator(builder, vmFileName),
		builder: builder,
		vmFileName: vmFileName,
	}
}

func (context *CodeContext) TranslateArithmetic(commandName string) {
	context.arithmeticTranslator.Translate(commandName)
}

func (context *CodeContext) TranslateMemory(pushOrPop string, segment string, index string) {
	context.memorySegmentTranslator.Translate(pushOrPop, segment, index)
}

func (context *CodeContext) GetCodeString() string {
	return context.builder.String()
}

// Get starting value from optional parameters of CodeContext. Needed for multi-threading
func getStartingBranchNumber(params ...interface{}) int {
	if len(params) > 2 {
		log.Fatal("Invalid arguments for CodeContext")
	} else if len(params) == 2 {
		return params[1].(int)
	}

	return 0
}
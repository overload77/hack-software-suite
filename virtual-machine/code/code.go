package code

import (
	"log"
	"strings"
)

type CodeContext struct {
	arithmeticTranslator *ArithmeticTranslator
	memorySegmentTranslator *MemorySegmentTranslator
	builder *strings.Builder
}

func GetCodeContext(params ...int) *CodeContext {
	startingBranchNum := getStartingBranchNumber(params...)
	builder := &strings.Builder{}
	return &CodeContext {
		arithmeticTranslator: GetArithmeticTranslator(startingBranchNum, builder),
		memorySegmentTranslator: GetMemorySegmentTranslator(builder),
		builder: builder,
	}
}

func (context *CodeContext) TranslateArithmetic(commandName string) {
	context.arithmeticTranslator.Handlers[commandName]()
}

func (context *CodeContext) TranslateMemory(pushOrPop string, segment string, index string) {
	context.memorySegmentTranslator.Handlers[pushOrPop](segment, index)
}

func (context *CodeContext) GetCodeString() string {
	return context.builder.String()
}

// Get starting value from optional parameters of CodeContext. Needed for multi-threading
func getStartingBranchNumber(params ...int) int {
	if len(params) > 1 {
		log.Fatal("Invalid arguments for CodeContext")
	} else if len(params) == 1 {
		return params[0]
	}

	return 0
}
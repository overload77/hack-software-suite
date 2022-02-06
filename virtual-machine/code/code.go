package code

import "strings"

type CodeContext struct {
	arithmeticHandler *ArithmeticCommand
	// memoryHandler *MemoryCommand
	builder strings.Builder
}

func GetCodeContext() *CodeContext {
	return &CodeContext {
		arithmeticHandler: GetArithmeticCommand(),
		builder: strings.Builder{},
	}
}

func (context *CodeContext) TranslateArithmetic(commandName string) {
	context.arithmeticHandler.Handlers[commandName](context.builder)
}

// TODO
func (context *CodeContext) TranslateMemory(pushOrPop string, segment string, index string) {
	context.builder.WriteString("todo")
}

func (context *CodeContext) GetCodeString() string {
	return context.builder.String()
}
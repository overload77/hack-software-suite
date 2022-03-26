package function

import (
	"strings"
	"github.com/overload77/hack-software-suite/virtual-machine/code/function/handlers"
)

type FunctionTranslator struct {
	callHandler *handlers.CallHandler
	functionHandler *handlers.CallHandler
	returnHandler *handlers.CallHandler
	currentFunction string
	builder *strings.Builder
}

func GetFunctionTranslator(builder *strings.Builder, vmFileName string) *FunctionTranslator {
	return &FunctionTranslator {
		callHandler: handlers.GetCallHandler(builder),
		currentFunction: "default",
	}
}

func (translator *FunctionTranslator) Translate(command, firstArg, secondArg string) {
	switch command {
	case "call":
		translator.callHandler.HandleTranslation(translator.currentFunction, firstArg, secondArg)
	case "function": // FROM HERE
		// GetCallHandler(dispatcher.builder)
	case "return":		
		// GetCallHandler(dispatcher.builder)
	}
	
	// GetCallHandler(dispatcher.builder) // Temp
}
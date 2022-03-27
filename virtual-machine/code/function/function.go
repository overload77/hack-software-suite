package function

import (
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code/function/handlers"
)

type FunctionTranslator struct {
	callHandler *handlers.CallHandler
	declarationHandler *handlers.DeclarationHandler
	returnHandler *handlers.CallHandler
	currentFunction *string
	builder *strings.Builder
}

func GetFunctionTranslator(builder *strings.Builder, currentFunction *string) *FunctionTranslator {
	return &FunctionTranslator {
		callHandler: handlers.GetCallHandler(builder),
		declarationHandler: handlers.GetDeclarationHandler(builder),
		currentFunction: currentFunction,
	}
}

func (translator *FunctionTranslator) Translate(command, firstArg, secondArg string) {
	switch command {
	case "call":
		translator.callHandler.HandleTranslation(*translator.currentFunction, firstArg, secondArg)
	case "function":
		*translator.currentFunction = firstArg
		translator.declarationHandler.HandleTranslation(firstArg, secondArg)
	case "return":		
		// GetCallHandler(dispatcher.builder)
	}
	
	// GetCallHandler(dispatcher.builder) // Temp
}
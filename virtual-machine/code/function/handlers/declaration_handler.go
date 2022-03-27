package handlers

import (
	"fmt"
	"strconv"
	"strings"
)

type DeclarationHandler struct {
	builder *strings.Builder
	functionName string
	nVars int
}

func GetDeclarationHandler(builder *strings.Builder) *DeclarationHandler {
	return &DeclarationHandler{builder: builder}
}

func (handler *DeclarationHandler) HandleTranslation(functionName, nVars string) {
	handler.updateArguments(functionName, nVars)
	handler.commentThisCall()
	handler.injectLabel()
	handler.initLocalVariables()
}

func (handler *DeclarationHandler) updateArguments(functionName, nVars string) {
	nVarsAsInt, _ := strconv.Atoi(nVars)
	handler.functionName = functionName
	handler.nVars = nVarsAsInt
}

func (handler *DeclarationHandler) commentThisCall() {
	handler.builder.WriteString(fmt.Sprintf("// Function %s %d\n", handler.functionName, handler.nVars))
}

func (handler *DeclarationHandler) injectLabel() {
	handler.builder.WriteString(fmt.Sprintf("(%s)\n", handler.functionName))
}

func (handler *DeclarationHandler) initLocalVariables() {
	for i := 0; i < handler.nVars; i++ {
		handler.pushZeroToTheStack()
	}
}

func (handler *DeclarationHandler) pushZeroToTheStack() {
	handler.builder.WriteString("@SP\n")
	handler.builder.WriteString("M=M+1\n")
	handler.builder.WriteString("A=M-1\n")
	handler.builder.WriteString("M=0\n")
}
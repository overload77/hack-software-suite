package handlers

import (
	"fmt"
	"strconv"
	"strings"
)

type CallHandler struct {
	builder *strings.Builder
	functionToReturn string
	functionToCall string
	nArgs int
	returnAddrCounter map[string]int
}

func GetCallHandler(builder *strings.Builder) *CallHandler {
	return &CallHandler{
		builder: builder,
		returnAddrCounter: map[string]int{},
	}
}

func (handler *CallHandler) HandleTranslation(functionToReturn, functionToCall, nArgs string) {
	handler.updateArguments(functionToReturn, functionToCall, nArgs)
	returnAddress := handler.getReturnAddress()
	handler.commentThisCall(functionToCall, nArgs)
	handler.pushReturnAddress(returnAddress)
	handler.pushLocal()
	handler.pushArg()
	handler.pushThis()
	handler.pushThat()
	handler.repositionArg()
	handler.jumpToFunction()
	handler.plantReturnLabel(returnAddress)
}

func (handler *CallHandler) updateArguments(functionToReturn, functionName, nArgs string) {
	nArgsAsInt, _ := strconv.Atoi(nArgs)
	handler.functionToReturn = functionToReturn
	handler.functionToCall = functionName
	handler.nArgs = nArgsAsInt
}

func (handler *CallHandler) getReturnAddress() string {
	returnCount, isThere := handler.returnAddrCounter[handler.functionToReturn]
	if isThere {
		handler.returnAddrCounter[handler.functionToReturn]++
	} else {
		handler.returnAddrCounter[handler.functionToReturn] = 1
	}
	return fmt.Sprintf("%s$ret.%d", handler.functionToReturn, returnCount)
}

func (handler *CallHandler) commentThisCall(functionToCall, nArgs string) {
	handler.builder.WriteString(fmt.Sprintf("// Call %s %s\n", functionToCall, nArgs))
}

func (handler *CallHandler) pushReturnAddress(returnAddress string) {
	handler.pushSymbolToTheStack(returnAddress)
}

func (handler *CallHandler) pushLocal() {
	handler.pushSegmentToTheStack("LCL")
}

func (handler *CallHandler) pushArg() {
	handler.pushSegmentToTheStack("ARG")
}

func (handler *CallHandler) pushThis() {
	handler.pushSegmentToTheStack("THIS")
}

func (handler *CallHandler) pushThat() {
	handler.pushSegmentToTheStack("THAT")
}

func (handler *CallHandler) repositionArg() {
	offset := 5 + handler.nArgs
	handler.builder.WriteString(fmt.Sprintf("@%d\n", offset))
	handler.builder.WriteString("D=A\n")
	handler.builder.WriteString("@SP\n")
	handler.builder.WriteString("D=M-D\n")
	handler.builder.WriteString("@ARG\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *CallHandler) jumpToFunction() {
	handler.builder.WriteString(fmt.Sprintf("@%s\n", handler.functionToCall))
	handler.builder.WriteString("0;JMP\n")
}

func (handler *CallHandler) plantReturnLabel(returnAddress string) {
	handler.builder.WriteString(fmt.Sprintf("(%s)\n", returnAddress))
}

// Little duplicated code but not worth to refactor
func (handler *CallHandler) pushSymbolToTheStack(symbol string) {
	handler.builder.WriteString(fmt.Sprintf("@%s\n", symbol))
	handler.builder.WriteString("D=A\n")
	handler.pushDRegisterToTheStack()
}

func (handler *CallHandler) pushSegmentToTheStack(segment string) {
	handler.builder.WriteString(fmt.Sprintf("@%s\n", segment))
	handler.builder.WriteString("D=M\n")
	handler.pushDRegisterToTheStack()
}

func (handler *CallHandler) pushDRegisterToTheStack() {
	handler.builder.WriteString("@SP\n")
	handler.builder.WriteString("M=M+1\n")
	handler.builder.WriteString("A=M-1\n")
	handler.builder.WriteString("M=D\n")
}

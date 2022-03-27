package handlers

import "strings"

type ReturnHandler struct {
	builder *strings.Builder
}

func GetReturnHandler(builder *strings.Builder) *ReturnHandler {
	return &ReturnHandler{builder: builder}
}

func (handler *ReturnHandler) HandleTranslation() {
	handler.commentThisCall()
	handler.saveLocalToR13()
	handler.saveReturnAddrToR14()
	handler.popReturnValueToARG()
	handler.repositionSP()
	handler.restoreTHAT()
	handler.restoreTHIS()
	handler.restoreARG()
	handler.restoreLCL()
	handler.gotoReturnAddr()
}

func (handler *ReturnHandler) commentThisCall() {
	handler.builder.WriteString("// Return\n")
}

func (handler *ReturnHandler) saveLocalToR13() {
	handler.builder.WriteString("@LCL\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@R13\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *ReturnHandler) saveReturnAddrToR14() {
	handler.builder.WriteString("@5\n")
	handler.builder.WriteString("A=D-A\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@R14\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *ReturnHandler) popReturnValueToARG() {
	handler.builder.WriteString("@SP\n")
	handler.builder.WriteString("A=M-1\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@ARG\n")
	handler.builder.WriteString("A=M\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *ReturnHandler) repositionSP() {
	handler.builder.WriteString("@ARG\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@SP\n")
	handler.builder.WriteString("M=D+1\n")
}

func (handler *ReturnHandler) restoreTHAT() {
	handler.builder.WriteString("@R13\n")
	handler.builder.WriteString("A=M-1\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@THAT\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *ReturnHandler) restoreTHIS() {
	handler.builder.WriteString("@R13\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@2\n")
	handler.builder.WriteString("A=D-A\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@THIS\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *ReturnHandler) restoreARG() {
	handler.builder.WriteString("@R13\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@3\n")
	handler.builder.WriteString("A=D-A\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@ARG\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *ReturnHandler) restoreLCL() {
	handler.builder.WriteString("@R13\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@4\n")
	handler.builder.WriteString("A=D-A\n")
	handler.builder.WriteString("D=M\n")
	handler.builder.WriteString("@LCL\n")
	handler.builder.WriteString("M=D\n")
}

func (handler *ReturnHandler) gotoReturnAddr() {
	handler.builder.WriteString("@R14\n")
	handler.builder.WriteString("A=M\n")
	handler.builder.WriteString("0;JMP\n")
}
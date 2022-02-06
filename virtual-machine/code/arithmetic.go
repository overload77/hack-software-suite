package code

import (
	"fmt"
	"strings"
)

type ArithmeticCommand struct {
	Handlers map[string]func(*strings.Builder)
	currentBranchNum int
}


func GetArithmeticCommand(startingBranchNum int) *ArithmeticCommand {
	arithmeticCommand := &ArithmeticCommand{currentBranchNum: startingBranchNum}
	arithmeticCommand.Handlers = map[string]func(*strings.Builder) {
		"add": (*arithmeticCommand).translateAdd,
		"sub": (*arithmeticCommand).translateSub,
		"neg": (*arithmeticCommand).translateNeg,
		"eq": (*arithmeticCommand).translateEq,
		"gt": (*arithmeticCommand).translateGt,
		"lt": (*arithmeticCommand).translateLt,
		"and": (*arithmeticCommand).translateAnd,
		"or": (*arithmeticCommand).translateOr,
		"not": (*arithmeticCommand).translateNot,
	}
	return arithmeticCommand
}


func (arithmeticCommand *ArithmeticCommand) translateAdd(builder *strings.Builder) {
	builder.WriteString("// Add\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=D+M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateSub(builder *strings.Builder) {
	builder.WriteString("// Sub\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=M-D\n")
}

func (arithmeticCommand *ArithmeticCommand) translateNeg(builder *strings.Builder) {
	builder.WriteString("// Neg\n")
	builder.WriteString("@SP\n")
	builder.WriteString("A=M-1\n")
	builder.WriteString("M=-M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateEq(builder *strings.Builder) {
	builder.WriteString("// Eq\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("D=M-D\n")
	builder.WriteString("M=1\n")
	builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("D;JNE\n")
	builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("0;JMP\n")
	builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("M=0\n")
	builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.currentBranchNum++
}

func (arithmeticCommand *ArithmeticCommand) translateGt(builder *strings.Builder) {
	builder.WriteString("// Gt\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("D=M-D\n")
	builder.WriteString("M=1\n")
	builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("D;JLE\n")
	builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("0;JMP\n")
	builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("M=0\n")
	builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.currentBranchNum++
}

func (arithmeticCommand *ArithmeticCommand) translateLt(builder *strings.Builder) {
	builder.WriteString("// Lt\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("D=M-D\n")
	builder.WriteString("M=1\n")
	builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("D;JGE\n")
	builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("0;JMP\n")
	builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticCommand.currentBranchNum))
	builder.WriteString("M=0\n")
	builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.currentBranchNum++
}

func (arithmeticCommand *ArithmeticCommand) translateAnd(builder *strings.Builder) {
	builder.WriteString("// And\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=D&M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateOr(builder *strings.Builder) {
	builder.WriteString("// Or\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=D|M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateNot(builder *strings.Builder) {
	builder.WriteString("// Not\n")
	builder.WriteString("@SP\n")
	builder.WriteString("A=M-1\n")
	builder.WriteString("M=!M\n")
}
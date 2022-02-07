package code

import (
	"fmt"
	"log"
	"strings"
)

type ArithmeticCommand struct {
	Handlers map[string]func()
	currentBranchNum int
	builder *strings.Builder
}


func GetArithmeticCommand(members ...interface{}) *ArithmeticCommand {
	startingBranchNum, builder := getBranchNumAndBuilder(members...)
	arithmeticCommand := &ArithmeticCommand {
		currentBranchNum: startingBranchNum,
		builder: builder,
	}
	arithmeticCommand.Handlers = map[string]func() {
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

// Helper function to interpret and return branch number and builder from variable parameters
func getBranchNumAndBuilder(members ...interface{}) (int, *strings.Builder) {
	if memberLen := len(members); memberLen == 2 {
		return members[0].(int), members[1].(*strings.Builder)
	} else if memberLen != 0 {
		log.Fatal("Wrong initialization of ArithmeticCommand")
	}
	
	return 0, &strings.Builder{}
}


func (arithmeticCommand *ArithmeticCommand) translateAdd() {
	arithmeticCommand.builder.WriteString("// Add\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("M=M-1\n")
	arithmeticCommand.builder.WriteString("A=M\n")
	arithmeticCommand.builder.WriteString("D=M\n")
	arithmeticCommand.builder.WriteString("A=A-1\n")
	arithmeticCommand.builder.WriteString("M=D+M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateSub() {
	arithmeticCommand.builder.WriteString("// Sub\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("M=M-1\n")
	arithmeticCommand.builder.WriteString("A=M\n")
	arithmeticCommand.builder.WriteString("D=M\n")
	arithmeticCommand.builder.WriteString("A=A-1\n")
	arithmeticCommand.builder.WriteString("M=M-D\n")
}

func (arithmeticCommand *ArithmeticCommand) translateNeg() {
	arithmeticCommand.builder.WriteString("// Neg\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("A=M-1\n")
	arithmeticCommand.builder.WriteString("M=-M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateEq() {
	arithmeticCommand.builder.WriteString("// Eq\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("M=M-1\n")
	arithmeticCommand.builder.WriteString("A=M\n")
	arithmeticCommand.builder.WriteString("D=M\n")
	arithmeticCommand.builder.WriteString("A=A-1\n")
	arithmeticCommand.builder.WriteString("D=M-D\n")
	arithmeticCommand.builder.WriteString("M=1\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("D;JNE\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("0;JMP\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("M=0\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.currentBranchNum++
}

func (arithmeticCommand *ArithmeticCommand) translateGt() {
	arithmeticCommand.builder.WriteString("// Gt\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("M=M-1\n")
	arithmeticCommand.builder.WriteString("A=M\n")
	arithmeticCommand.builder.WriteString("D=M\n")
	arithmeticCommand.builder.WriteString("A=A-1\n")
	arithmeticCommand.builder.WriteString("D=M-D\n")
	arithmeticCommand.builder.WriteString("M=1\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("D;JLE\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("0;JMP\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("M=0\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.currentBranchNum++
}

func (arithmeticCommand *ArithmeticCommand) translateLt() {
	arithmeticCommand.builder.WriteString("// Lt\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("M=M-1\n")
	arithmeticCommand.builder.WriteString("A=M\n")
	arithmeticCommand.builder.WriteString("D=M\n")
	arithmeticCommand.builder.WriteString("A=A-1\n")
	arithmeticCommand.builder.WriteString("D=M-D\n")
	arithmeticCommand.builder.WriteString("M=1\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("D;JGE\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("0;JMP\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.builder.WriteString("M=0\n")
	arithmeticCommand.builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticCommand.currentBranchNum))
	arithmeticCommand.currentBranchNum++
}

func (arithmeticCommand *ArithmeticCommand) translateAnd() {
	arithmeticCommand.builder.WriteString("// And\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("M=M-1\n")
	arithmeticCommand.builder.WriteString("A=M\n")
	arithmeticCommand.builder.WriteString("D=M\n")
	arithmeticCommand.builder.WriteString("A=A-1\n")
	arithmeticCommand.builder.WriteString("M=D&M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateOr() {
	arithmeticCommand.builder.WriteString("// Or\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("M=M-1\n")
	arithmeticCommand.builder.WriteString("A=M\n")
	arithmeticCommand.builder.WriteString("D=M\n")
	arithmeticCommand.builder.WriteString("A=A-1\n")
	arithmeticCommand.builder.WriteString("M=D|M\n")
}

func (arithmeticCommand *ArithmeticCommand) translateNot() {
	arithmeticCommand.builder.WriteString("// Not\n")
	arithmeticCommand.builder.WriteString("@SP\n")
	arithmeticCommand.builder.WriteString("A=M-1\n")
	arithmeticCommand.builder.WriteString("M=!M\n")
}
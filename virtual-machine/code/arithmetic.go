package code

import (
	"fmt"
	"log"
	"strings"
)

type ArithmeticTranslator struct {
	Handlers map[string]func()
	currentBranchNum int
	builder *strings.Builder
}

func GetArithmeticTranslator(members ...interface{}) *ArithmeticTranslator {
	startingBranchNum, builder := getBranchNumAndBuilder(members...)
	arithmeticTranslator := &ArithmeticTranslator {
		currentBranchNum: startingBranchNum,
		builder: builder,
	}
	arithmeticTranslator.Handlers = map[string]func() {
		"add": arithmeticTranslator.translateAdd,
		"sub": arithmeticTranslator.translateSub,
		"neg": arithmeticTranslator.translateNeg,
		"eq": arithmeticTranslator.translateEq,
		"gt": arithmeticTranslator.translateGt,
		"lt": arithmeticTranslator.translateLt,
		"and": arithmeticTranslator.translateAnd,
		"or": arithmeticTranslator.translateOr,
		"not": arithmeticTranslator.translateNot,
	}
	
	return arithmeticTranslator
}

func (arithmeticTranslator *ArithmeticTranslator) Translate(command string) {
	handlerMethod, isOk := arithmeticTranslator.Handlers[command]
	if !isOk {
		log.Fatalln("Invalid arithmetic command")
	}
	handlerMethod()
}

// Helper function to interpret and return branch number and builder from variable parameters
func getBranchNumAndBuilder(members ...interface{}) (int, *strings.Builder) {
	if memberLen := len(members); memberLen == 2 {
		return members[0].(int), members[1].(*strings.Builder)
	} else if memberLen != 0 {
		log.Fatal("Wrong initialization of ArithmeticTranslator")
	}
	
	return 0, &strings.Builder{}
}


func (arithmeticTranslator *ArithmeticTranslator) translateAdd() {
	arithmeticTranslator.builder.WriteString("// Add\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("M=M-1\n")
	arithmeticTranslator.builder.WriteString("A=M\n")
	arithmeticTranslator.builder.WriteString("D=M\n")
	arithmeticTranslator.builder.WriteString("A=A-1\n")
	arithmeticTranslator.builder.WriteString("M=D+M\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateSub() {
	arithmeticTranslator.builder.WriteString("// Sub\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("M=M-1\n")
	arithmeticTranslator.builder.WriteString("A=M\n")
	arithmeticTranslator.builder.WriteString("D=M\n")
	arithmeticTranslator.builder.WriteString("A=A-1\n")
	arithmeticTranslator.builder.WriteString("M=M-D\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateNeg() {
	arithmeticTranslator.builder.WriteString("// Neg\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("A=M-1\n")
	arithmeticTranslator.builder.WriteString("M=-M\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateEq() {
	arithmeticTranslator.builder.WriteString("// Eq\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("M=M-1\n")
	arithmeticTranslator.builder.WriteString("A=M\n")
	arithmeticTranslator.builder.WriteString("D=M\n")
	arithmeticTranslator.builder.WriteString("A=A-1\n")
	arithmeticTranslator.builder.WriteString("D=M-D\n")
	arithmeticTranslator.builder.WriteString("M=-1\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("D;JNE\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("0;JMP\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("A=M-1\n")
	arithmeticTranslator.builder.WriteString("M=0\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.currentBranchNum++
}

func (arithmeticTranslator *ArithmeticTranslator) translateGt() {
	arithmeticTranslator.builder.WriteString("// Gt\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("M=M-1\n")
	arithmeticTranslator.builder.WriteString("A=M\n")
	arithmeticTranslator.builder.WriteString("D=M\n")
	arithmeticTranslator.builder.WriteString("A=A-1\n")
	arithmeticTranslator.builder.WriteString("D=M-D\n")
	arithmeticTranslator.builder.WriteString("M=-1\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("D;JLE\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("0;JMP\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("A=M-1\n")
	arithmeticTranslator.builder.WriteString("M=0\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.currentBranchNum++
}

func (arithmeticTranslator *ArithmeticTranslator) translateLt() {
	arithmeticTranslator.builder.WriteString("// Lt\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("M=M-1\n")
	arithmeticTranslator.builder.WriteString("A=M\n")
	arithmeticTranslator.builder.WriteString("D=M\n")
	arithmeticTranslator.builder.WriteString("A=A-1\n")
	arithmeticTranslator.builder.WriteString("D=M-D\n")
	arithmeticTranslator.builder.WriteString("M=-1\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("@ELSE%d\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("D;JGE\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("0;JMP\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("(ELSE%d)\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("A=M-1\n")
	arithmeticTranslator.builder.WriteString("M=0\n")
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.currentBranchNum++
}

func (arithmeticTranslator *ArithmeticTranslator) translateAnd() {
	arithmeticTranslator.builder.WriteString("// And\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("M=M-1\n")
	arithmeticTranslator.builder.WriteString("A=M\n")
	arithmeticTranslator.builder.WriteString("D=M\n")
	arithmeticTranslator.builder.WriteString("A=A-1\n")
	arithmeticTranslator.builder.WriteString("M=D&M\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateOr() {
	arithmeticTranslator.builder.WriteString("// Or\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("M=M-1\n")
	arithmeticTranslator.builder.WriteString("A=M\n")
	arithmeticTranslator.builder.WriteString("D=M\n")
	arithmeticTranslator.builder.WriteString("A=A-1\n")
	arithmeticTranslator.builder.WriteString("M=D|M\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateNot() {
	arithmeticTranslator.builder.WriteString("// Not\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("A=M-1\n")
	arithmeticTranslator.builder.WriteString("M=!M\n")
}
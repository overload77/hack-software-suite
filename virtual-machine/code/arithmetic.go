package code

import (
	"fmt"
	"log"
	"strings"
)

type ArithmeticTranslator struct {
	Handlers map[string]func()
	builder *strings.Builder
	currentBranchNum int
	vmFileName string
}

func GetArithmeticTranslator(builder *strings.Builder, vmFileName string) *ArithmeticTranslator {
	arithmeticTranslator := &ArithmeticTranslator {
		builder: builder,
		currentBranchNum: 0,
		vmFileName: vmFileName,
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

func (arithmeticTranslator *ArithmeticTranslator) translateAdd() {
	arithmeticTranslator.builder.WriteString("// Add\n")
	popStackToDandLoadNextValueToM(arithmeticTranslator.builder)
	arithmeticTranslator.builder.WriteString("M=D+M\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateSub() {
	arithmeticTranslator.builder.WriteString("// Sub\n")
	popStackToDandLoadNextValueToM(arithmeticTranslator.builder)
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
	popStackToDandLoadNextValueToM(arithmeticTranslator.builder)
	writeComparisonCommands(arithmeticTranslator, "JNE")
}

func (arithmeticTranslator *ArithmeticTranslator) translateGt() {
	arithmeticTranslator.builder.WriteString("// Gt\n")
	popStackToDandLoadNextValueToM(arithmeticTranslator.builder)
	writeComparisonCommands(arithmeticTranslator, "JLE")
}

func (arithmeticTranslator *ArithmeticTranslator) translateLt() {
	arithmeticTranslator.builder.WriteString("// Lt\n")
	popStackToDandLoadNextValueToM(arithmeticTranslator.builder)
	writeComparisonCommands(arithmeticTranslator, "JGE")
}

func (arithmeticTranslator *ArithmeticTranslator) translateAnd() {
	arithmeticTranslator.builder.WriteString("// And\n")
	popStackToDandLoadNextValueToM(arithmeticTranslator.builder)
	arithmeticTranslator.builder.WriteString("M=D&M\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateOr() {
	arithmeticTranslator.builder.WriteString("// Or\n")
	popStackToDandLoadNextValueToM(arithmeticTranslator.builder)
	arithmeticTranslator.builder.WriteString("M=D|M\n")
}

func (arithmeticTranslator *ArithmeticTranslator) translateNot() {
	arithmeticTranslator.builder.WriteString("// Not\n")
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("A=M-1\n")
	arithmeticTranslator.builder.WriteString("M=!M\n")
}

// Pop the stack into D register. Then decrement A by one to set M the next value in the stack
func popStackToDandLoadNextValueToM(builder *strings.Builder) {
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
}

// Translate EQ, GT and LT commands. jumpInstr is JNE for eq, JLE for gt and JGT for lt
func writeComparisonCommands(arithmeticTranslator *ArithmeticTranslator, jumpInstr string) {
	arithmeticTranslator.builder.WriteString("D=M-D\n")
	arithmeticTranslator.builder.WriteString("M=-1\n")
	arithmeticTranslator.builder.WriteString(
		fmt.Sprintf("@ELSE.%s.%d\n", arithmeticTranslator.vmFileName,
					arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString(fmt.Sprintf("D;%s\n", jumpInstr))
	arithmeticTranslator.builder.WriteString(
		fmt.Sprintf("@CONTINUE.%s.%d\n", arithmeticTranslator.vmFileName,
					arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("0;JMP\n")
	arithmeticTranslator.builder.WriteString(
		fmt.Sprintf("(ELSE.%s.%d)\n", arithmeticTranslator.vmFileName,
					arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.builder.WriteString("@SP\n")
	arithmeticTranslator.builder.WriteString("A=M-1\n")
	arithmeticTranslator.builder.WriteString("M=0\n")
	arithmeticTranslator.builder.WriteString(
		fmt.Sprintf("(@CONTINUE.%s.%d)\n", arithmeticTranslator.vmFileName,
					arithmeticTranslator.currentBranchNum))
	arithmeticTranslator.currentBranchNum++
}

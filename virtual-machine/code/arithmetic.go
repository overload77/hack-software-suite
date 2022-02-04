package code

import (
	"fmt"
	"strings"
)

var currentBranchNum = 0
var ArithmeticCommandsAndWriters = map[string]func() string {
	"add": translateAdd,
	"sub": translateSub,
	"neg": translateNeg,
	"eq": translateWithBranch(translateEq),
	"gt": translateWithBranch(translateGt),
	"lt": translateWithBranch(translateLt),
	"and": translateAnd,
	"or": translateOr,
	"not": translateNot,
}

// Closure for eq, gt and lt commands. Increments branch number
func translateWithBranch(handler func() string) func() string {
	return func() string {
		currentBranchNum++
		return handler()
	}
}

func translateAdd() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Add\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("M=M-1\n")
	asmLine.WriteString("A=M\n")
	asmLine.WriteString("D=M\n")
	asmLine.WriteString("A=A-1\n")
	asmLine.WriteString("M=D+M\n")
	return asmLine.String()
}

func translateSub() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Sub\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("M=M-1\n")
	asmLine.WriteString("A=M\n")
	asmLine.WriteString("D=M\n")
	asmLine.WriteString("A=A-1\n")
	asmLine.WriteString("M=M-D\n")
	return asmLine.String()
}

func translateNeg() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Neg\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("A=M-1\n")
	asmLine.WriteString("M=-M\n")
	return asmLine.String()
}

func translateEq() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Eq\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("M=M-1\n")
	asmLine.WriteString("A=M\n")
	asmLine.WriteString("D=M\n")
	asmLine.WriteString("A=A-1\n")
	asmLine.WriteString("D=M-D\n")
	asmLine.WriteString("M=1\n")
	asmLine.WriteString(fmt.Sprintf("@ELSE%d\n", currentBranchNum))
	asmLine.WriteString("D;JNE\n")
	asmLine.WriteString(fmt.Sprintf("@CONTINUE%d\n", currentBranchNum))
	asmLine.WriteString("0;JMP\n")
	asmLine.WriteString(fmt.Sprintf("(ELSE%d)\n", currentBranchNum))
	asmLine.WriteString("M=0\n")
	asmLine.WriteString(fmt.Sprintf("(CONTINUE%d)\n", currentBranchNum))
	return asmLine.String()
}

func translateGt() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Gt\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("M=M-1\n")
	asmLine.WriteString("A=M\n")
	asmLine.WriteString("D=M\n")
	asmLine.WriteString("A=A-1\n")
	asmLine.WriteString("D=M-D\n")
	asmLine.WriteString("M=1\n")
	asmLine.WriteString(fmt.Sprintf("@ELSE%d\n", currentBranchNum))
	asmLine.WriteString("D;JLE\n")
	asmLine.WriteString(fmt.Sprintf("@CONTINUE%d\n", currentBranchNum))
	asmLine.WriteString("0;JMP\n")
	asmLine.WriteString(fmt.Sprintf("(ELSE%d)\n", currentBranchNum))
	asmLine.WriteString("M=0\n")
	asmLine.WriteString(fmt.Sprintf("(CONTINUE%d)\n", currentBranchNum))
	return asmLine.String()
}

func translateLt() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Lt\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("M=M-1\n")
	asmLine.WriteString("A=M\n")
	asmLine.WriteString("D=M\n")
	asmLine.WriteString("A=A-1\n")
	asmLine.WriteString("D=M-D\n")
	asmLine.WriteString("M=1\n")
	asmLine.WriteString(fmt.Sprintf("@ELSE%d\n", currentBranchNum))
	asmLine.WriteString("D;JGE\n")
	asmLine.WriteString(fmt.Sprintf("@CONTINUE%d\n", currentBranchNum))
	asmLine.WriteString("0;JMP\n")
	asmLine.WriteString(fmt.Sprintf("(ELSE%d)\n", currentBranchNum))
	asmLine.WriteString("M=0\n")
	asmLine.WriteString(fmt.Sprintf("(CONTINUE%d)\n", currentBranchNum))
	return asmLine.String()
}

func translateAnd() string {
	var asmLine strings.Builder
	asmLine.WriteString("// And\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("M=M-1\n")
	asmLine.WriteString("A=M\n")
	asmLine.WriteString("D=M\n")
	asmLine.WriteString("A=A-1\n")
	asmLine.WriteString("M=D&M\n")
	return asmLine.String()
}

func translateOr() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Or\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("M=M-1\n")
	asmLine.WriteString("A=M\n")
	asmLine.WriteString("D=M\n")
	asmLine.WriteString("A=A-1\n")
	asmLine.WriteString("M=D|M\n")
	return asmLine.String()
}

func translateNot() string {
	var asmLine strings.Builder
	asmLine.WriteString("// Not\n")
	asmLine.WriteString("@SP\n")
	asmLine.WriteString("A=M-1\n")
	asmLine.WriteString("M=!M\n")
	return asmLine.String()
}
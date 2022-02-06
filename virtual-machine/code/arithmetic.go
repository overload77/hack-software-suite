package code

import (
	"fmt"
	"strings"
)

var ArithmeticCommandsAndWriters = map[string]interface{} {
	"add": translateAdd,
	"sub": translateSub,
	"neg": translateNeg,
	"eq": translateEq,
	"gt": translateGt,
	"lt": translateLt,
	"and": translateAnd,
	"or": translateOr,
	"not": translateNot,
}


func translateAdd(builder strings.Builder) strings.Builder {
	builder.WriteString("// Add\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=D+M\n")
	return builder
}

func translateSub(builder strings.Builder) strings.Builder {
	builder.WriteString("// Sub\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=M-D\n")
	return builder
}

func translateNeg(builder strings.Builder) strings.Builder {
	builder.WriteString("// Neg\n")
	builder.WriteString("@SP\n")
	builder.WriteString("A=M-1\n")
	builder.WriteString("M=-M\n")
	return builder
}

func translateEq(builder strings.Builder, branchNum int) strings.Builder {
	builder.WriteString("// Eq\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("D=M-D\n")
	builder.WriteString("M=1\n")
	builder.WriteString(fmt.Sprintf("@ELSE%d\n", branchNum))
	builder.WriteString("D;JNE\n")
	builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", branchNum))
	builder.WriteString("0;JMP\n")
	builder.WriteString(fmt.Sprintf("(ELSE%d)\n", branchNum))
	builder.WriteString("M=0\n")
	builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", branchNum))
	return builder
}

func translateGt(builder strings.Builder, branchNum int) strings.Builder {
	builder.WriteString("// Gt\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("D=M-D\n")
	builder.WriteString("M=1\n")
	builder.WriteString(fmt.Sprintf("@ELSE%d\n", branchNum))
	builder.WriteString("D;JLE\n")
	builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", branchNum))
	builder.WriteString("0;JMP\n")
	builder.WriteString(fmt.Sprintf("(ELSE%d)\n", branchNum))
	builder.WriteString("M=0\n")
	builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", branchNum))
	return builder
}

func translateLt(builder strings.Builder, branchNum int) strings.Builder {
	builder.WriteString("// Lt\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("D=M-D\n")
	builder.WriteString("M=1\n")
	builder.WriteString(fmt.Sprintf("@ELSE%d\n", branchNum))
	builder.WriteString("D;JGE\n")
	builder.WriteString(fmt.Sprintf("@CONTINUE%d\n", branchNum))
	builder.WriteString("0;JMP\n")
	builder.WriteString(fmt.Sprintf("(ELSE%d)\n", branchNum))
	builder.WriteString("M=0\n")
	builder.WriteString(fmt.Sprintf("(CONTINUE%d)\n", branchNum))
	return builder
}

func translateAnd(builder strings.Builder) strings.Builder {
	builder.WriteString("// And\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=D&M\n")
	return builder
}

func translateOr(builder strings.Builder) strings.Builder {
	builder.WriteString("// Or\n")
	builder.WriteString("@SP\n")
	builder.WriteString("M=M-1\n")
	builder.WriteString("A=M\n")
	builder.WriteString("D=M\n")
	builder.WriteString("A=A-1\n")
	builder.WriteString("M=D|M\n")
	return builder
}

func translateNot(builder strings.Builder) strings.Builder {
	builder.WriteString("// Not\n")
	builder.WriteString("@SP\n")
	builder.WriteString("A=M-1\n")
	builder.WriteString("M=!M\n")
	return builder
}
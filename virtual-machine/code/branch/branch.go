package branch

import (
	"fmt"
	"log"
	"strings"
)

type BranchTranslator struct {
	Handlers map[string]func(string)
	builder *strings.Builder
	vmFileName string
	currentFunction string
}

func GetBranchTranslator(builder *strings.Builder, vmFileName string) *BranchTranslator {
	branchTranslator := &BranchTranslator {
		builder: builder,
		vmFileName: vmFileName,
		currentFunction: "null",
	}
	branchTranslator.Handlers = map[string]func(string) {
		"label": branchTranslator.translateLabel,
		"goto": branchTranslator.translateGoto,
		"if-goto": branchTranslator.translateIfGoto,
	}
	
	return branchTranslator
}

func (translator *BranchTranslator) Translate(command, label, dummy string) {
	if handlerMethod, isOk := translator.Handlers[command]; isOk {
		handlerMethod(label)
	} else {
		log.Fatalln("Invalid branching command")
	}
}

func (translator *BranchTranslator) translateLabel(label string) {
	translator.builder.WriteString(fmt.Sprintf("// Label %s\n", label))
	label = fmt.Sprintf("(%s.%s$%s)\n", translator.vmFileName,
						translator.currentFunction, label)
	translator.builder.WriteString(label)
}

func (translator *BranchTranslator) translateGoto(label string) {
	translator.builder.WriteString(fmt.Sprintf("// Goto %s\n", label))
	jumpLocation := fmt.Sprintf("@%s.%s$%s\n", translator.vmFileName,
								translator.currentFunction, label)
	translator.builder.WriteString(jumpLocation)
	translator.builder.WriteString("0;JMP\n")
}

func (translator *BranchTranslator) translateIfGoto(label string) {
	translator.builder.WriteString(fmt.Sprintf("// If-Goto %s\n", label))
	jumpLocation := fmt.Sprintf("@%s.%s$%s\n", translator.vmFileName,
								translator.currentFunction, label)
	translator.builder.WriteString("@SP\n")
	translator.builder.WriteString("M=M-1\n")
	translator.builder.WriteString("A=M\n")
	translator.builder.WriteString("D=M\n")
	translator.builder.WriteString(jumpLocation)
	translator.builder.WriteString("D;JNE\n")
}
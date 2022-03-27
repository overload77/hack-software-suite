package code

import (
	"bufio"
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code/arithmetic"
	"github.com/overload77/hack-software-suite/virtual-machine/code/branch"
	"github.com/overload77/hack-software-suite/virtual-machine/code/function"
	"github.com/overload77/hack-software-suite/virtual-machine/code/memory"
)

type CodeContext struct {
	arithmeticTranslator Translator
	memorySegmentTranslator Translator
	branchTranslator Translator
	functionTranslator Translator
	builder *strings.Builder
	vmFileName string
	currentTranslator Translator
	currentCommand string
	currentFirstArg string
	currentSecondArg string
	currentFunction *string
}

type Translator interface {
	Translate(string, string, string)
}

func GetCodeContext(vmFileName string) *CodeContext {
	builder := &strings.Builder{}
	currentFunction := "default"
	return &CodeContext {
		arithmeticTranslator: arithmetic.GetArithmeticTranslator(builder, vmFileName),
		memorySegmentTranslator: memory.GetMemorySegmentTranslator(builder, vmFileName),
		branchTranslator: branch.GetBranchTranslator(builder, &currentFunction),
		functionTranslator: function.GetFunctionTranslator(builder, &currentFunction),
		builder: builder,
		vmFileName: vmFileName,
		currentFunction: &currentFunction,
	}
}

func (context *CodeContext) TranslateCommand(command string) {
	if command, isValidCommand := trimLine(command); isValidCommand {
		context.parseLine(command)
		context.currentTranslator.Translate(
			context.currentCommand, context.currentFirstArg, context.currentSecondArg)
	}
}

func AddBootstrapCode(buffer *bufio.Writer) {
	buffer.WriteString("// Bootstrapping\n")
	buffer.WriteString("@256\n")
	buffer.WriteString("D=A\n")
	buffer.WriteString("@SP\n")
	buffer.WriteString("M=D\n")
	buffer.WriteString("@Sys.init\n")
	buffer.WriteString("0;JMP\n")
}

func (context *CodeContext) GetCodeString() string {
	return context.builder.String()
}

func trimLine(line string) (string, bool) {
	line = strings.Trim(line, " ")
	if strings.HasPrefix(line, "//") || len(line) == 0 {
		return "", false
	} else if commentStart := strings.Index(line, "//"); commentStart != -1 {
		return strings.Trim(line[:commentStart], " "), true
	}

	return line, true
}
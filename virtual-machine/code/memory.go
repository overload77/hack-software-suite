package code

import "fmt"
import "strings"

type MemorySegmentTranslator struct {
	Handlers map[string]func(string, string)
	builder *strings.Builder
	mappings map[string]string
}

func GetMemorySegmentTranslator(builder *strings.Builder) *MemorySegmentTranslator {
	memSegCommand := &MemorySegmentTranslator {
		builder: builder,
		mappings: map[string]string {
			"local": "LCL",
			"argument": "ARG",
			"this": "THIS",
			"that": "THAT",
		},
	}
	memSegCommand.Handlers = map[string]func(string, string) {
		"push": memSegCommand.translatePush,
		"pop": memSegCommand.translatePop,
	}

	return memSegCommand
}

func (memSegCommand *MemorySegmentTranslator) translatePush(segment string, index string) {
	segmentSymbol := memSegCommand.mappings[segment]
	memSegCommand.builder.WriteString(fmt.Sprintf("// Push %s %s\n", segment, index))
	memSegCommand.builder.WriteString(fmt.Sprintf("@%s\n", segmentSymbol))
	memSegCommand.builder.WriteString("D=M\n")
	memSegCommand.builder.WriteString(fmt.Sprintf("@%s\n", index))
	memSegCommand.builder.WriteString("A=D+A\n")
	memSegCommand.builder.WriteString("D=M\n")
	memSegCommand.builder.WriteString("@SP\n")
	memSegCommand.builder.WriteString("M=M+1\n")
	memSegCommand.builder.WriteString("A=M-1\n")
	memSegCommand.builder.WriteString("M=D\n")
}

func (memSegCommand *MemorySegmentTranslator) translatePop(segment string, index string) {
	segmentSymbol := memSegCommand.mappings[segment]
	memSegCommand.builder.WriteString(fmt.Sprintf("// Pop %s %s\n", segment, index))
	memSegCommand.builder.WriteString("@SP\n")
	memSegCommand.builder.WriteString("M=M-1\n")
	memSegCommand.builder.WriteString("A=M\n")
	memSegCommand.builder.WriteString("D=M\n")
	memSegCommand.builder.WriteString("@13\n")
	memSegCommand.builder.WriteString("M=D\n")

	memSegCommand.builder.WriteString(fmt.Sprintf("@%s\n", segmentSymbol))
	memSegCommand.builder.WriteString("D=M\n")
	memSegCommand.builder.WriteString(fmt.Sprintf("@%s\n", index))
	memSegCommand.builder.WriteString("D=D+A\n")
	memSegCommand.builder.WriteString("@14\n")
	memSegCommand.builder.WriteString("M=D\n")

	memSegCommand.builder.WriteString("@13\n")
	memSegCommand.builder.WriteString("D=M\n")
	memSegCommand.builder.WriteString("@14\n")
	memSegCommand.builder.WriteString("A=M\n")
	memSegCommand.builder.WriteString("M=D\n")
}
package memory

import (
	"fmt"
	"strings"
)

type MemorySegmentTranslator struct {
	handlers map[string]func(string, string)
	builder *strings.Builder
	currentSegment Segment
	vmFileName string
}

func GetMemorySegmentTranslator(builder *strings.Builder,
		vmFileName string) *MemorySegmentTranslator {
	memoryTranslator := &MemorySegmentTranslator {
		builder: builder,
		vmFileName: vmFileName,
	}
	memoryTranslator.handlers = map[string]func(string, string) {
		"push": memoryTranslator.translatePush,
		"pop": memoryTranslator.translatePop,
	}

	return memoryTranslator
}

func (memoryTranslator *MemorySegmentTranslator) Translate(command, segmentName, index string) {
	memoryTranslator.currentSegment = SegmentFactory(
		segmentName, index, memoryTranslator.vmFileName, memoryTranslator.builder)
	memoryTranslator.handlers[command](segmentName, index)
}

func (memoryTranslator *MemorySegmentTranslator) translatePush(segmentName, index string) {
	memoryTranslator.builder.WriteString(fmt.Sprintf("// Push %s %s\n", segmentName, index))
	memoryTranslator.currentSegment.writeValueToDReg()
	memoryTranslator.pushDRegisterToStack()
}

func (memoryTranslator *MemorySegmentTranslator) translatePop(segmentName, index string) {
	memoryTranslator.builder.WriteString(fmt.Sprintf("// Pop %s %s\n", segmentName, index))
	memoryTranslator.popStackToDReg()
	memoryTranslator.currentSegment.setFromDReg()
}

// Pushes D register to stack
func (memoryTranslator *MemorySegmentTranslator) pushDRegisterToStack() {
	memoryTranslator.builder.WriteString("@SP\n")
	memoryTranslator.builder.WriteString("M=M+1\n")
	memoryTranslator.builder.WriteString("A=M-1\n")
	memoryTranslator.builder.WriteString("M=D\n")
}

// Pops the stack and puts value in D register
func (memoryTranslator *MemorySegmentTranslator) popStackToDReg() {
	memoryTranslator.builder.WriteString("@SP\n")
	memoryTranslator.builder.WriteString("M=M-1\n")
	memoryTranslator.builder.WriteString("A=M\n")
	memoryTranslator.builder.WriteString("D=M\n")
}

package code

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type MemorySegmentTranslator struct {
	Handlers map[string]func(string, string)
	builder *strings.Builder
	mappings map[string]string
	vmFileName string
}

func GetMemorySegmentTranslator(builder *strings.Builder) *MemorySegmentTranslator {
	memoryTranslator := &MemorySegmentTranslator {
		builder: builder,
		mappings: map[string]string {
			"local": "LCL",
			"argument": "ARG",
			"this": "THIS",
			"that": "THAT",
		},
	}
	memoryTranslator.Handlers = map[string]func(string, string) {
		"push": memoryTranslator.translatePush,
		"pop": memoryTranslator.translatePop,
	}

	return memoryTranslator
}

func (memoryTranslator *MemorySegmentTranslator) translatePush(segment, index string) {
	memoryTranslator.builder.WriteString(fmt.Sprintf("// Push %s %s\n", segment, index))
	memoryTranslator.writeSegmentValueToDReg(segment, index)
	memoryTranslator.pushDRegisterToStack()
}

func (memoryTranslator *MemorySegmentTranslator) translatePop(segment, index string) {
	memoryTranslator.builder.WriteString(fmt.Sprintf("// Pop %s %s\n", segment, index))
	memoryTranslator.popStackToDReg()
	memoryTranslator.pushDRegisterToSegment(segment, index)
}

// Get's segments value and puts it into D register
func (memoryTranslator *MemorySegmentTranslator) writeSegmentValueToDReg(segment, index string) {
	switch segment {
	case "pointer":
		memoryTranslator.writePointerToDReg(index)
	case "temp":
		memoryTranslator.writeTempToDReg(index)
	case "constant":
		memoryTranslator.writeConstantToDReg(index)
	case "static":
		memoryTranslator.writeStaticToDReg(index)
	default:
		memoryTranslator.writeMappedSegmentsToDReg(segment, index)
	}
}

// Writes first 4 segments(LCL, ARG, THIS and THAT) value to the D register
func (memoryTranslator *MemorySegmentTranslator) writeMappedSegmentsToDReg(segment, index string) {
	segmentSymbol, ok := memoryTranslator.mappings[segment]
	if !ok {
		log.Fatal("Invalid memory segment")
	}
	memoryTranslator.builder.WriteString(fmt.Sprintf("@%s\n", segmentSymbol))
	memoryTranslator.builder.WriteString("D=M\n")
	memoryTranslator.builder.WriteString(fmt.Sprintf("@%s\n", index))
	memoryTranslator.builder.WriteString("A=D+A\n")
	memoryTranslator.builder.WriteString("D=M\n")
}

// Writes pointer segment to D register
func (memoryTranslator *MemorySegmentTranslator) writePointerToDReg(index string) {
	if index == "0" {
		memoryTranslator.builder.WriteString("@THIS\n")
	} else if index == "1" {
		memoryTranslator.builder.WriteString("@THAT\n")
	} else {
		log.Fatal("Invalid pointer segment index")
	}
	memoryTranslator.builder.WriteString("D=M\n")
}

// Writes temp segment to D register
func (memoryTranslator *MemorySegmentTranslator) writeTempToDReg(index string) {
	integerIndex, _ := strconv.Atoi(index)
	tempRegister := fmt.Sprintf("%d", 5 + integerIndex)
	memoryTranslator.builder.WriteString(fmt.Sprintf("@%s\n", tempRegister))
	memoryTranslator.builder.WriteString("D=M\n")
}

// Writes constant segment to D register
func (memoryTranslator *MemorySegmentTranslator) writeConstantToDReg(index string) {
	memoryTranslator.builder.WriteString(fmt.Sprintf("@%s\n", index))
	memoryTranslator.builder.WriteString("D=A\n")
}

// Writes static segment to D register
func (memoryTranslator *MemorySegmentTranslator) writeStaticToDReg(index string) {
	staticVarSymbol := fmt.Sprintf("%s.%s", memoryTranslator.vmFileName, index)
	memoryTranslator.builder.WriteString(fmt.Sprintf("@%s\n", staticVarSymbol))
	memoryTranslator.builder.WriteString("D=M\n")
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

// Writes value in D register to the segment
func (memoryTranslator *MemorySegmentTranslator) pushDRegisterToSegment(segment, index string) {
	switch segment {
	case "pointer":
		memoryTranslator.writeDRegToPointer(index)
	default:
		memoryTranslator.writeDRegToMappedSegments(segment, index)
	}
}

// Writes popped value to mapped segents(LCL, ARG, THIS and THAT)
// Stores popped value and pointer to the segment entry temporarily in registers 13 & 14
func (memoryTranslator *MemorySegmentTranslator) writeDRegToMappedSegments(segment, index string) {
	memoryTranslator.writeDRegToR13()
	memoryTranslator.writeSegmentPointerToR14(segment, index)
	memoryTranslator.writeR13ToSegmentStoredOnR14()
}

// Writes D register to the pointer
func (memoryTranslator *MemorySegmentTranslator) writeDRegToPointer(index string) {
	if index == "0" {
		memoryTranslator.builder.WriteString("@THIS\n")
	} else if index == "1" {
		memoryTranslator.builder.WriteString("@THAT\n")
	} else {
		log.Fatal("Invalid pointer segment index")
	}
	memoryTranslator.builder.WriteString("M=D\n")
}

// Puts value in D register to register 13
func (memoryTranslator *MemorySegmentTranslator) writeDRegToR13() {
	memoryTranslator.builder.WriteString("@13\n")
	memoryTranslator.builder.WriteString("M=D\n")
}

// Put pointer to the segment to register 14
func (memoryTranslator *MemorySegmentTranslator) writeSegmentPointerToR14(segment, index string) {
	segmentSymbol, ok := memoryTranslator.mappings[segment]
	if !ok {
		log.Fatal("Invalid memory segment")
	}
	memoryTranslator.builder.WriteString(fmt.Sprintf("@%s\n", segmentSymbol))
	memoryTranslator.builder.WriteString("D=M\n")
	memoryTranslator.builder.WriteString(fmt.Sprintf("@%s\n", index))
	memoryTranslator.builder.WriteString("D=D+A\n")
	memoryTranslator.builder.WriteString("@14\n")
	memoryTranslator.builder.WriteString("M=D\n")
}

// Put value in register 13(Popped value from the stack) to the segment
// which it's pointer is on register 14
func (memoryTranslator *MemorySegmentTranslator) writeR13ToSegmentStoredOnR14() {
	memoryTranslator.builder.WriteString("@13\n")
	memoryTranslator.builder.WriteString("D=M\n")
	memoryTranslator.builder.WriteString("@14\n")
	memoryTranslator.builder.WriteString("A=M\n")
	memoryTranslator.builder.WriteString("M=D\n")
}

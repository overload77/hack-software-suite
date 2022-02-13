package memory

import (
	"fmt"
	"strings"
)

type MappedSegment struct {
	segmentSymbol string
	index string
	builder *strings.Builder
}

// Writes mapped(first 4) segments(LCL, ARG, THIS and THAT) value to the D register
func (mappedSegment *MappedSegment) writeValueToDReg() {
	mappedSegment.builder.WriteString(fmt.Sprintf("@%s\n", mappedSegment.segmentSymbol))
	mappedSegment.builder.WriteString("D=M\n")
	mappedSegment.builder.WriteString(fmt.Sprintf("@%s\n", mappedSegment.index))
	mappedSegment.builder.WriteString("A=D+A\n")
	mappedSegment.builder.WriteString("D=M\n")
}

// Writes popped value to mapped segments(LCL, ARG, THIS and THAT)
// Stores popped value and pointer to the segment entry temporarily in registers 13 & 14
func (mappedSegment *MappedSegment) setFromDReg() {
	mappedSegment.writeDRegToR13()
	mappedSegment.writeSegmentPointerToR14()
	mappedSegment.writeR13ToSegmentStoredOnR14()
}

// Puts value in D register to register 13
func (mappedSegment *MappedSegment) writeDRegToR13() {
	mappedSegment.builder.WriteString("@13\n")
	mappedSegment.builder.WriteString("M=D\n")
}

// Put pointer to the segment to register 14
func (mappedSegment *MappedSegment) writeSegmentPointerToR14() {
	mappedSegment.builder.WriteString(fmt.Sprintf("@%s\n", mappedSegment.segmentSymbol))
	mappedSegment.builder.WriteString("D=M\n")
	mappedSegment.builder.WriteString(fmt.Sprintf("@%s\n", mappedSegment.index))
	mappedSegment.builder.WriteString("D=D+A\n")
	mappedSegment.builder.WriteString("@14\n")
	mappedSegment.builder.WriteString("M=D\n")
}

// Put value in register 13(Popped value from the stack) to the segment
// which it's pointer is on register 14
func (mappedSegment *MappedSegment) writeR13ToSegmentStoredOnR14() {
	mappedSegment.builder.WriteString("@13\n")
	mappedSegment.builder.WriteString("D=M\n")
	mappedSegment.builder.WriteString("@14\n")
	mappedSegment.builder.WriteString("A=M\n")
	mappedSegment.builder.WriteString("M=D\n")
}
package memory

import (
	"fmt"
	"strings"
)

type StaticSegment struct {
	index string
	vmFileName string
	builder *strings.Builder
}

// Writes static segment to D register
func (staticSegment *StaticSegment) writeValueToDReg() {
	staticVarSymbol := fmt.Sprintf("%s.%s", staticSegment.vmFileName, staticSegment.index)
	staticSegment.builder.WriteString(fmt.Sprintf("@%s\n", staticVarSymbol))
	staticSegment.builder.WriteString("D=M\n")
}

// Writes D register to the static segment
func (staticSegment *StaticSegment) setFromDReg() {
	staticVarSymbol := fmt.Sprintf("%s.%s", staticSegment.vmFileName, staticSegment.index)
	staticSegment.builder.WriteString(fmt.Sprintf("@%s\n", staticVarSymbol))
	staticSegment.builder.WriteString("M=D\n")
}
package memory

import (
	"fmt"
	"strings"
)

type ConstantSegment struct {
	index string
	builder *strings.Builder
}

// Writes constant segment to D register
func (constantSegment *ConstantSegment) writeValueToDReg() {
	constantSegment.builder.WriteString(fmt.Sprintf("@%s\n", constantSegment.index))
	constantSegment.builder.WriteString("D=A\n")
}

// Dummy 
func (constantSegment *ConstantSegment) setFromDReg() {}
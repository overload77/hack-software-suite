package memory

import (
	"fmt"
	"strconv"
	"strings"
)

type TempSegment struct {
	index string
	builder *strings.Builder
}

// Writes temp segment to D register
func (tempSegment *TempSegment) writeValueToDReg() {
	integerIndex, _ := strconv.Atoi(tempSegment.index)
	tempRegister := fmt.Sprintf("%d", 5 + integerIndex)
	tempSegment.builder.WriteString(fmt.Sprintf("@%s\n", tempRegister))
	tempSegment.builder.WriteString("D=M\n")
}

// Writes D register to the temp segment
func (tempSegment *TempSegment) setFromDReg() {
	integerIndex, _ := strconv.Atoi(tempSegment.index)
	tempRegister := fmt.Sprintf("%d", 5 + integerIndex)
	tempSegment.builder.WriteString(fmt.Sprintf("@%s\n", tempRegister))
	tempSegment.builder.WriteString("M=D\n")
}
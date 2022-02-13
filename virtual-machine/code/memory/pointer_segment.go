package memory

import (
	"log"
	"strings"
)

type PointerSegment struct {
	index string
	builder *strings.Builder
}

// Writes pointer segment to D register
func (pointerSegment *PointerSegment) writeValueToDReg() {
	if pointerSegment.index == "0" {
		pointerSegment.builder.WriteString("@THIS\n")
	} else if pointerSegment.index == "1" {
		pointerSegment.builder.WriteString("@THAT\n")
	} else {
		log.Fatal("Invalid pointer segment index")
	}
	pointerSegment.builder.WriteString("D=M\n")
}

// Writes D register to the pointer segment
func (pointerSegment *PointerSegment) setFromDReg() {
	if pointerSegment.index == "0" {
		pointerSegment.builder.WriteString("@THIS\n")
	} else if pointerSegment.index == "1" {
		pointerSegment.builder.WriteString("@THAT\n")
	} else {
		log.Fatal("Invalid pointer segment index")
	}
	pointerSegment.builder.WriteString("M=D\n")
}
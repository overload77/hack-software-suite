// Segment interface and factory. Declares write to/from register methods
package memory

import (
	"log"
	"strings"
)

type Segment interface {
	writeValueToDReg()
	setFromDReg()
}

func SegmentFactory(segmentName, index, vmFileName string, builder *strings.Builder) Segment {
	switch segmentName {
	case "pointer":
		return &PointerSegment{index: index, builder: builder}
	case "temp":
		return &TempSegment{index: index, builder: builder}
	case "constant":
		return &ConstantSegment{index: index, builder: builder}
	case "static":
		return &StaticSegment{index: index, vmFileName: vmFileName, builder: builder}
	default:
		segmentSymbol := getSegmentMapping(segmentName)
		return &MappedSegment {segmentSymbol: segmentSymbol, index: index, builder: builder}
	}
}

// Gets the mappings of first 4 memory segments
func getSegmentMapping(segmentName string) string {
	mapping := map[string]string {
		"local": "LCL",
		"argument": "ARG",
		"this": "THIS",
		"that": "THAT",
	}
	segmentSymbol, isOk := mapping[segmentName]
	if !isOk {
		log.Fatalln("Invalid memory segment %s", segmentName)
	}
	return segmentSymbol
}

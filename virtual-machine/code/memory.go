package code

import "strings"

type MemorySegmentCommand struct {
	Handlers map[string]func(string, int)
	builder *strings.Builder
	mappings map[string]string
}

func GetMemorySegmentCommand(builder *strings.Builder) *MemorySegmentCommand {
	memSegCommand := &MemorySegmentCommand {
		builder: builder,
		mappings: map[string]string {
			"local": "LCL",
			"argument": "ARG",
			"this": "THIS",
			"that": "THAT",
		},
	}
	memSegCommand.Handlers = map[string]func(string, int) {
		"push": (*memSegCommand).translatePush,
		// "pop": (*memSegCommand).translatePop,
	}
	return memSegCommand
}

func (memSegCommand *MemorySegmentCommand) translatePush(segment string, index int) {

}
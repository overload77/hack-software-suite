package main

import (
	"fmt"
	"strings"
	"strconv"
)

func convertLine(line string) string {
	var binaryInstruction string
	if strings.HasPrefix(line, "@") {
		binaryInstruction = convertTypeA(line)
	} else {
		binaryInstruction = convertTypeC(line)
	}

	return binaryInstruction
}

// Returns 16 bit binary representation of A-type instruction 
func convertTypeA(line string) string {
	// TODO: Parser and C type instruction decoder before symbols
	valueOrSymbol := line[1:]
	value, _ := getValueOfTypeA(valueOrSymbol)
	spacePrependedRepr := fmt.Sprintf("%16b", value)
	return strings.ReplaceAll(spacePrependedRepr, " ", "0")
}


func convertTypeC(line string) string {
	return ""
}

// Gets integer value from value part of A-type instruction(After @ part)
// E.g: Returns 123 for @123, returns 16 for @firstVar
func getValueOfTypeA(valueOrSymbol string) (int, error) {
	value, err := strconv.Atoi(valueOrSymbol)
	if err != nil {
		// TODO: Lookup symbol table and get it
		fmt.Println("Error:", err)
	}
	return value, nil
}
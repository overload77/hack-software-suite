package code

import (
	"fmt"
	"strings"
	"strconv"
)

// Returns binary representation of instruction. Entrypoint for code package
// Determines instruction type(A or C) and passes them to correct handlers
func ConvertLine(line string) string {
	if strings.HasPrefix(line, "@") {
		return convertTypeA(line)
	} else {
		return convertTypeC(line)
	}
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
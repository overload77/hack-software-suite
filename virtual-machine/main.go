package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code"
	"github.com/overload77/hack-software-suite/virtual-machine/parser"
)

func main() {
	log.Println("Starting")
	validateArgument()
	
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		commandType, commandName, firstArg, secondArg := parser.ParseLine(line)
		fmt.Println("Parsed line:", commandType, commandName, firstArg, secondArg)
		// coded := code.TranslateCommand(commandType, commandName, firstArg, secondArg)
		// fmt.Println("Coded:", coded)
	}

	file.Close()
}

func validateArgument() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	} else if !strings.HasSuffix(os.Args[1], ".vm") {
		log.Fatal("Invalid file extension. Should end with .vm")
	}
}

func translateCommand(commandType parser.CommandType, commandName string,
					  firstCommandArg string, secondCommandArg string) string {
	if commandType == parser.Arithmetic {
		return code.TranslateArithmetic(commandName)
	} else if commandType == parser.Memory {
		return code.TranslateMemory(commandName, firstCommandArg, secondCommandArg)
	}

	return ""
}

// Returns source and output files
func openFiles(filename string) (*os.File, *os.File) {
	sourceFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	asmFile, err := os.Create(strings.ReplaceAll(filename, ".vm", ".asm"))
	if err != nil {
		log.Fatal(err)
	}

	return sourceFile, asmFile
}
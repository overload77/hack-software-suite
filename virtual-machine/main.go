package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code"
	"github.com/overload77/hack-software-suite/virtual-machine/parser"
)

func main() {
	validateArgument()
	sourceFile, asmFile := openFiles(os.Args[1])
	translate(sourceFile, asmFile, os.Args[1])
	closeFiles(sourceFile, asmFile)
	log.Println("Done!")
}

func validateArgument() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	} else if !strings.HasSuffix(os.Args[1], ".vm") {
		log.Fatal("Invalid file extension. Should end with .vm")
	}
}

func translate(sourceFile, asmFile *os.File, sourceFileName string) {
	scanner := bufio.NewScanner(sourceFile)
	writer := bufio.NewWriter(asmFile)
	codeContext := code.GetCodeContext(sourceFileName)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(line, "//") {
			continue
		}
		commandType, commandName, firstArg, secondArg := parser.ParseLine(line)
		translateSingleCommand(codeContext, commandType, commandName, firstArg, secondArg)
	}
	writer.WriteString(codeContext.GetCodeString() + "\n")
	writer.Flush()
}

func translateSingleCommand(codeContext *code.CodeContext, commandType parser.CommandType,
			commandName string, firstCommandArg string, secondCommandArg string) {
	if commandType == parser.Arithmetic {
		codeContext.TranslateArithmetic(commandName)
	} else if commandType == parser.Memory {
		codeContext.TranslateMemory(commandName, firstCommandArg, secondCommandArg)
	}
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

func closeFiles(sourceFile, asmFile *os.File) {
	sourceFile.Close()
	asmFile.Close()
}
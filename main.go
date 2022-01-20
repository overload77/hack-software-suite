package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/overload77/go-hack-assembler/code"
	"github.com/overload77/go-hack-assembler/instructionset"
	"github.com/overload77/go-hack-assembler/symboltable"
)


func main() {
	log.Println("Starting")
	if err := validateArgument(); err != nil {
		log.Fatal(err)
	}
	filename := os.Args[1]
	symbolTable := symboltable.NewSymbolTable()
	firstPass(filename, symbolTable)
	secondPass(filename, symbolTable)
}


func validateArgument() error {
	if len(os.Args) != 2 {
		return errors.New("Invalid number of arguments")
	} else if !strings.HasSuffix(os.Args[1], ".asm") {
		return errors.New("Invalid file extension. Should end with .asm")
	}

	return nil
}

func firstPass(filename string, symbolTable *symboltable.SymbolTable) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	currentInstructionAddr := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "//") || len(line) == 0 {
			continue
		} else if labelStart := strings.Index(line, "("); labelStart != -1 {
			symbol := line[labelStart + 1:strings.Index(line, ")")]
			symbolTable.AddLabelToTable(symbol, currentInstructionAddr)
			continue
		}
		currentInstructionAddr++
	}
}

func secondPass(filename string, symbolTable *symboltable.SymbolTable) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	hackFile, err := os.Create(strings.ReplaceAll(filename, ".asm", ".hack"))
	if err != nil {
		log.Fatal(err)
	}

	instructionset := instructionset.NewCInstructionSet()
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(hackFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "//") || strings.Contains(line, "(") || len(line) == 0 {
			continue
		}

		binaryInstr := code.ConvertLine(line, symbolTable, instructionset)
		writer.WriteString(binaryInstr + "\n")
	}
	writer.Flush()
	hackFile.Close()
}
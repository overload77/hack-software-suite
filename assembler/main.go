package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/overload77/hack-software-suite/assembler/code"
	"github.com/overload77/hack-software-suite/assembler/instructionset"
	"github.com/overload77/hack-software-suite/assembler/symboltable"
)

func main() {
	validateArgument()
	assemble()
	log.Println("Completed!")
}

func validateArgument() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	} else if !strings.HasSuffix(os.Args[1], ".asm") {
		log.Fatal("Invalid file extension. Should end with .asm")
	}
}

// Starts two-pass assembling process
func assemble() {
	filename := os.Args[1]
	symbolTable := symboltable.NewSymbolTable()
	runFirstPass(filename, symbolTable)
	runSecondPass(filename, symbolTable)
}

// First pass populates symbol table with labels
func runFirstPass(filename string, symbolTable *symboltable.SymbolTable) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	currentInstructionAddr := 0
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		if strings.HasPrefix(line, "//") || len(line) == 0 {
			continue
		} else if labelStart := strings.Index(line, "("); labelStart != -1 {
			symbol := line[labelStart + 1:strings.Index(line, ")")]
			symbolTable.AddLabelToTable(symbol, currentInstructionAddr)
			continue
		}
		currentInstructionAddr++
	}

	file.Close()
}

// Second pass converts instructions line by line and writes to file in ASCII format
func runSecondPass(filename string, symbolTable *symboltable.SymbolTable) {
	sourceFile, hackFile := openFiles(filename)
	instructionset := instructionset.NewCInstructionSet()
	
	scanner := bufio.NewScanner(sourceFile)
	writer := bufio.NewWriter(hackFile)
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		if strings.HasPrefix(line, "//") || strings.Contains(line, "(") || len(line) == 0 {
			continue
		}
		binaryInstr := code.ConvertLine(line, symbolTable, instructionset)
		writer.WriteString(binaryInstr + "\n")
	}

	writer.Flush()
	sourceFile.Close()
	hackFile.Close()
}

// Returns source and output files
func openFiles(filename string) (*os.File, *os.File) {
	sourceFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	hackFile, err := os.Create(strings.ReplaceAll(filename, ".asm", ".hack"))
	if err != nil {
		log.Fatal(err)
	}

	return sourceFile, hackFile
}
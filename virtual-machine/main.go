package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code"
)

func main() {
	validateArgument()
	sourceFiles, asmFile := openFiles(os.Args[1])
	translate(sourceFiles, asmFile)
	closeFiles(sourceFiles, asmFile)
	log.Println("Done!")
}

func validateArgument() {
	if len(os.Args) != 2 {
		log.Fatalln("Invalid number of arguments")
	} else if strings.HasSuffix(os.Args[1], "/") {
		log.Fatalln("Should not end with trailing slash")
	}
}

func translate(sourceFiles []map[string]interface{}, asmFile *os.File) {
	bootstrapStategy(sourceFiles[0], asmFile)
	for _, sourceFile := range sourceFiles {
		translateVmFile(sourceFile, asmFile)
	}
}

func translateVmFile(sourceFile map[string]interface{}, asmFile *os.File) {
	scanner := bufio.NewScanner(sourceFile["file"].(*os.File))
	writer := bufio.NewWriter(asmFile)
	codeContext := code.GetCodeContext(sourceFile["filename"].(string))
	for scanner.Scan() {
		codeContext.TranslateCommand(scanner.Text())
	}
	writer.WriteString(codeContext.GetCodeString())
	writer.Flush()
}

func bootstrapStategy(firstVmFile map[string]interface{}, asmFile *os.File) {
	if firstVmFile["filename"].(string) == "Sys.vm" {
		writer := bufio.NewWriter(asmFile)
		code.AddBootstrapCode(writer)
		writer.Flush()
	}
}
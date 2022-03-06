package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/overload77/hack-software-suite/virtual-machine/code"
	"github.com/overload77/hack-software-suite/virtual-machine/parser"
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
		log.Fatalln("Should not and with trailing slash")
	}
}

func translate(sourceFiles []map[string]interface{}, asmFile *os.File) {
	for _, sourceFile := range sourceFiles {
		translateVmFile(sourceFile, asmFile)
	}
}

func translateVmFile(sourceFile map[string]interface{}, asmFile *os.File) {
	scanner := bufio.NewScanner(sourceFile["file"].(*os.File))
	writer := bufio.NewWriter(asmFile)
	codeContext := code.GetCodeContext(sourceFile["filename"].(string))
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(line, "//") {
			continue
		}
		commandType, commandName, firstArg, secondArg := parser.ParseLine(line)
		codeContext.TranslateCommand(commandType, commandName, firstArg, secondArg)
	}
	writer.WriteString(codeContext.GetCodeString() + "\n")
	writer.Flush()
}

// Returns source and output files
func openFiles(dirname string) ([]map[string]interface{}, *os.File) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatalln("Invalid directory")
	}

	sourceFiles := []map[string]interface{}{}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".vm") {
			continue
		}
		sourceFile, err := os.Open(dirname + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		} 

		sourceFiles = append(sourceFiles, map[string]interface{}{
			"file": sourceFile,
			"filename": file.Name(),
		})
	}

	asmFile := openOutputFile(dirname)
	return sourceFiles, asmFile
}

func openOutputFile(dirname string) *os.File {
	asmFile, err := os.Create(dirname + "/" + dirname + ".asm")
	if err != nil {
		log.Fatal(err)
	}

	return asmFile
}

func closeFiles(sourceFiles []map[string]interface{}, asmFile *os.File) {
	for _, file := range sourceFiles {
		file["file"].(*os.File).Close()
	}
	asmFile.Close()
}
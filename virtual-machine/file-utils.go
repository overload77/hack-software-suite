// This file contains all functions related to source/output file management
package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Returns source and output files
func openFiles(dirOrFilename string) ([]map[string]interface{}, *os.File) {
	if strings.HasSuffix(dirOrFilename, ".vm") {
		asmFilename := strings.ReplaceAll(dirOrFilename, ".vm", ".asm")
		return openSingleSourceFile(dirOrFilename), openOutputFile(asmFilename)
	}

	asmFilename := dirOrFilename + "/" + dirOrFilename + ".asm"
	return openSourceFilesFromDir(dirOrFilename), openOutputFile(asmFilename)
}

func openSingleSourceFile(filename string) []map[string]interface{} {
	sourceFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	return []map[string]interface{} {
		{"file": sourceFile, "filename": filename},
	}
}

func openSourceFilesFromDir(dirname string) []map[string]interface{} {
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

		sourceFiles = append(sourceFiles, map[string]interface{} {
			"file": sourceFile,
			"filename": file.Name(),
		})
	}

	return sourceFiles
}

func openOutputFile(filepath string) *os.File {
	asmFile, err := os.Create(filepath)
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
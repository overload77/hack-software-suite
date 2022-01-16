package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/overload77/go-hack-assembler/code"
)

func main() {
	fmt.Println("Starting")
	if err := validateArgument(); err != nil {
		log.Fatal(err)
	}

	filename := os.Args[1]
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		fmt.Println(code.ConvertLine(line))
    }

}


func validateArgument() error {
	if len(os.Args) != 2 {
		return errors.New("Invalid number of arguments")
	} else if !strings.HasSuffix(os.Args[1], ".asm") {
		return errors.New("Invalid file extension. Should end with .asm")
	}

	return nil
}
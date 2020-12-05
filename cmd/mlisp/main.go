package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"bakku.dev/minimalisp"
	"github.com/peterh/liner"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: mlisp [script]")
	} else if len(os.Args) == 2 {
		runScript(os.Args[1])
	} else {
		startRepl()
	}
}

func runScript(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("could not read file: %v", err))
		return
	}

	code := string(bytes)

	scanner := minimalisp.NewScanner(code, os.Stdout)
	tokens, ok := scanner.Scan()

	if !ok {
		return
	}

	parser := minimalisp.NewParser(tokens)
	expressions, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	interpreter := minimalisp.NewInterpreter()
	ret, err := interpreter.Interpret(expressions)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("=> %v", ret))
	}
}

func startRepl() {
	interpreter := minimalisp.NewInterpreter()
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	for {
		code, err := line.Prompt("> ")
		if err == liner.ErrPromptAborted {
			continue
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		code = strings.TrimSpace(code)
		if code == "exit" {
			return
		}

		line.AppendHistory(code)

		scanner := minimalisp.NewScanner(code, os.Stdout)
		tokens, ok := scanner.Scan()
		if !ok {
			continue
		}

		parser := minimalisp.NewParser(tokens)
		expressions, err := parser.Parse()
		if err != nil {
			fmt.Println(err)
			continue
		}

		ret, err := interpreter.Interpret(expressions)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("=> %v", ret))
		}
	}
}

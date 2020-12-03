package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"bakku.dev/tinylisp"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: tl [script]")
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

	scanner := tinylisp.NewScanner(code, os.Stdout)
	tokens, err := scanner.Scan()

	if err == nil {
		fmt.Println(tokens)
	}
}

func startRepl() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		code, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		scanner := tinylisp.NewScanner(code, os.Stdout)
		tokens, err := scanner.Scan()

		if err == nil {
			fmt.Println(tokens)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runRepl(version string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to misery " + version)

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		text = strings.Trim(text, "\n")
		if len(text) < 1 {
			continue
		}

		if text == "exit" {
			return
		}

		lexer := NewLexer(text)
		parser := NewParser(lexer)

		trees, err := parser.Parse()

		if err != nil {
			fmt.Println(err)
		}

		for _, tree := range trees {
			fmt.Println(tree.String())
		}

	}
}

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		panic("Using files is not supported yet")
	}

	runRepl("0.1.0")

}

package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  rune = '('
	RIGHT_PAREN rune = ')'
	LEFT_BRACE  rune = '{'
	RIGHT_BRACE rune = '}'
	COMMA       rune = ','
	DOT         rune = '.'
	MINUS       rune = '-'
	PLUS        rune = '+'
	SEMICOLON   rune = ';'
	STAR        rune = '*'
	SLASH       rune = '/'
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	raw, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fileContents := string(raw)
	hasError := false
	lineNum := 1
	for _, token := range fileContents {
		switch token {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")
		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")
		case COMMA:
			fmt.Println("COMMA , null")
		case DOT:
			fmt.Println("DOT . null")
		case MINUS:
			fmt.Println("MINUS - null")
		case PLUS:
			fmt.Println("PLUS + null")
		case SEMICOLON:
			fmt.Println("SEMICOLON ; null")
		case STAR:
			fmt.Println("STAR * null")
		case SLASH:
			fmt.Println("SLASH / null")
		case '\n':
			lineNum++
		default:

			hasError = true
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", lineNum, string(token))
		}
	}
	//eof = end of file
	fmt.Println("EOF  null")
	if hasError {
		os.Exit(65)
	}
}

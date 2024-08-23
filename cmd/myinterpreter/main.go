package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	EQUAL       rune = '='
	BANG        rune = '!'
	LESS        rune = '<'
	GREATER     rune = '>'
	STRING      rune = '"'
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
	skip := false
	skipLine := false
	//might be easier to change to a for loop (easier to skip lines lol)
	for i := 0; i < len(fileContents); i++ {
		token := rune(fileContents[i])
		if skip || skipLine {
			skip = false
			if token == '\n' {
				skipLine = false
			}
		} else {
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
				if len(fileContents) > i+1 {
					if fileContents[i+1] == '/' {
						skip = true
						skipLine = true
						lineNum++
					} else {
						fmt.Println("SLASH / null")
					}
				} else {
					fmt.Println("SLASH / null")
				}
			case EQUAL:
				if len(fileContents) > i+1 {
					if fileContents[i+1] == '=' {
						fmt.Println("EQUAL_EQUAL == null")
						skip = true
					} else {
						fmt.Println("EQUAL = null")
					}
				} else {
					fmt.Println("EQUAL = null")
				}
			case BANG:
				if len(fileContents) > i+1 {
					if fileContents[i+1] == '=' {
						fmt.Println("BANG_EQUAL != null")
						skip = true
					} else {
						fmt.Println("BANG ! null")
					}
				} else {
					fmt.Println("BANG ! null")
				}
			case LESS:
				if len(fileContents) > i+1 {
					if fileContents[i+1] == '=' {
						fmt.Println("LESS_EQUAL <= null")
						skip = true
					} else {
						fmt.Println("LESS < null")
					}
				} else {
					fmt.Println("LESS < null")
				}
			case GREATER:
				if len(fileContents) > i+1 {
					if fileContents[i+1] == '=' {
						fmt.Println("GREATER_EQUAL >= null")
						skip = true
					} else {
						fmt.Println("GREATER > null")
					}
				} else {
					fmt.Println("GREATER > null")
				}
			case STRING:
				if len(fileContents) < i+1 {
					fmt.Println("[line 1] Error: Unterminated string")
				} else {
					it := i + 1
					isString := false
					for it < len(fileContents) {
						if fileContents[it] == '"' {
							var out strings.Builder
							isString = true
							for j := i + 1; j < it; j++ {
								out.WriteString(string(fileContents[j]))
							}
							i += it - i
							fmt.Printf("STRING \"%s\" %s\n", out.String(), out.String())
							break
						}
						it++
					}
					if !isString {
						i += it - i
						hasError = true
						fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", lineNum)
					}

				}
			case ' ', '\t':
				//skip
			case '\n':
				lineNum++
			default:
				if isDigit(token) {
					var out strings.Builder
					for i+1 < len(fileContents) && isDigit(rune(fileContents[i+1])) {
						out.WriteString(string(fileContents[i]))
						i++
					}
					if i+1 < len(fileContents) && fileContents[i+1] == '.' {
						out.WriteString(string(fileContents[i]))
						if i+2 < len(fileContents) && isDigit(rune(fileContents[i+2])) {
							out.WriteString(string(fileContents[i+1]))
							i += 2
							for i < len(fileContents) && isDigit(rune(fileContents[i])) {
								out.WriteString(string(fileContents[i]))
								i++
							}
						}
					} else if i < len(fileContents) && fileContents[i] != '.' {
						out.WriteString(string(fileContents[i]))
					}
					foutput, err := strconv.ParseFloat(out.String(), 64)
					if err != nil {
						hasError = true
						fmt.Fprintf(os.Stderr, "[line %d] Error: Invalid number: %s\n", lineNum, foutput)
						continue
					} else {
						foutput := fmt.Sprintf("%.6f", foutput)
						foutput = strings.TrimRight(foutput, "0")
						if foutput[len(foutput)-1] == '.' {
							foutput += "0"
						}
						fmt.Println("NUMBER", out.String(), foutput)
					}
					if i < len(fileContents) && fileContents[i] == '.' {
						fmt.Println("DOT . null")
					}
				} else {
					hasError = true
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", lineNum, string(token))
				}
			}
		}
	}
	//eof = end of file
	fmt.Println("EOF  null")
	if hasError {
		os.Exit(65)
	}
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

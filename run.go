package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to read the content of a file
func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func help() {
	fmt.Println("Type 'help' to show this text.")
	fmt.Println("Type 'quit' to exit.")
	fmt.Println("Type 'run <code>' to run some Brain**** code.")
	fmt.Println("Type 'file <file>' to run a Brain**** file.")
}

// Main function
func main() {
	fmt.Println("Brain**** Interpreter")
	help()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "exit":
			fallthrough
		case "quit":
			fmt.Println("Exiting program...")
			return
		case "help":
			help()
		default:
			if len(command) > 4 && command[:4] == "run " {
				code := command[4:]
				BFInterpreter(code)
			} else if len(command) > 5 && command[:5] == "file " {
				fileName := command[5:]
				//fmt.Println(fileName[len(fileName)-3:])
				if fileName[len(fileName)-3:] != ".bf" {
					fileName += ".bf"
				}
				//fmt.Println(fileName)
				content, err := readFile(fileName)
				if err != nil {
					fmt.Printf("Error reading file: %v\n", err)
				} else {
					BFInterpreter(content)
				}
			} else {
				fmt.Println("Command not recognized.")
			}
		}
	}
}

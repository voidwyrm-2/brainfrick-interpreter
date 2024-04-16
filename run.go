package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func writeFile(filename string, data string) error {
	// Open the file with write permissions, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func help() {
	fmt.Println("Type 'help' to show this text.")
	fmt.Println("type 'bfhelp' to show the instructions for Brain****")
	fmt.Println("Type 'quit' to exit.")
	fmt.Println("Type 'run <code>' to run some Brain**** code.")
	fmt.Println("Type 'file <file>' to run a Brain**** file.")
	fmt.Println("Type 'compress <file>' to remove all non-action characters in a Brain**** file, result written.")
}

func bfhelp() {
	fmt.Println("'>': Increment the data pointer by one (to point to the next cell to the right).")
	fmt.Println("'<': Decrement the data pointer by one (to point to the next cell to the left).")
	fmt.Println("'+': Increment the byte at the data pointer by one.")
	fmt.Println("'-': Decrement the byte at the data pointer by one.")
	fmt.Println("'.': Output the byte at the data pointer.")
	fmt.Println("',': Accept one byte of input, storing its value in the byte at the data pointer.")
	fmt.Println("'[': If the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching ']' command.")
	fmt.Println("']': If the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching '[' command.")
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
		case "bfhelp":
			bfhelp()
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
			} else if len(command) > 5 && command[:9] == "compress " {
				fileName := command[9:]
				//fmt.Println(fileName[len(fileName)-3:])
				if fileName[len(fileName)-3:] != ".bf" {
					fileName += ".bf"
				}
				//fmt.Println(fileName)
				content, err := readFile(fileName)
				if err != nil {
					fmt.Printf("Error reading file: %v\n", err)
				} else {
					compressed := BFCompressor(content)
					werr := writeFile(fileName[:3]+"_compressed.bf", compressed)
					if werr != nil {
						fmt.Printf("Error writing file: %v\n", werr)
					}
				}
			} else {
				fmt.Printf("Command '%s' not recognized.\n", strings.Split(command, " ")[0])
			}
		}
	}
}

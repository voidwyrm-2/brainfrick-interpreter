package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BFInterpreter(code string) {
	var bytes [30000]byte
	for i := range bytes {
		bytes[i] = 0
	}
	var pointer int = 0
	chars := strings.Split(code, "")

	jumps := make(map[int]int)
	var jumpstack []int
	var nests = 0
	for i, c := range chars {
		if c == "[" {
			nests += 1
			jumpstack = append(jumpstack, i)
		} else if c == "]" && nests > 0 {
			opening := jumpstack[len(jumpstack)-1]
			jumps[opening] = i
			jumps[i] = opening
			jumpstack = jumpstack[:len(jumpstack)-1]
			nests -= 1
		}
	}

	var charindex int = 0
	for charindex < len(chars) {

		c := chars[charindex]
		//fmt.Println(charindex, c)
		switch c {
		case "+":
			bytes[pointer] += 1
			charindex += 1
		case "-":
			bytes[pointer] -= 1
			charindex += 1
		case ">":
			pointer += 1
			charindex += 1
		case "<":
			pointer -= 1
			charindex += 1
		case ".":
			fmt.Println(string(bytes[pointer]))
			charindex += 1
		case ",":
			fmt.Println("please input a number")
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print(":: ")
			scanner.Scan()
			gotten := scanner.Text()
			if gotten != "--STOP" {
				bytes[pointer] = byte(gotten[0])
			}
			charindex += 1
		case "[":
			if bytes[pointer] == 0 {
				charindex = jumps[charindex] + 1
			} else {
				charindex += 1
			}
		case "]":
			if bytes[pointer] != 0 {
				charindex = jumps[charindex] + 1
			} else {
				charindex += 1
			}
		default:
			charindex += 1
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			if bytes[pointer]+1 > 255 {
				bytes[pointer] = 0
			} else {
				bytes[pointer] += 1
			}
			charindex += 1
		case "-":
			if int(bytes[pointer])-1 < 0 {
				bytes[pointer] = 255
			} else {
				bytes[pointer] -= 1
			}
			charindex += 1
		case ">":
			if pointer+1 > 29999 {
				pointer = 0
			} else {
				pointer += 1
			}
			charindex += 1
		case "<":
			if pointer-1 < 0 {
				pointer = 29999
			} else {
				pointer -= 1
			}
			charindex += 1
		case ".":
			fmt.Printf("b%v: %v('%s')\n", pointer, int(bytes[pointer]), string(bytes[pointer]))
			charindex += 1
		case ",":
			fmt.Println("please input a character(or do '-h' for other options)")
			scanner := bufio.NewScanner(os.Stdin)
			for {
				fmt.Print(":: ")
				scanner.Scan()
				gotten := scanner.Text()
				if gotten == "-h" {
					fmt.Println("do '-n <decimal>' to input a decimal literal")
					fmt.Println("do '-b <binary>' to input a binary literal")
					fmt.Println("do '-x <hex>' to input a hexidecimal literal")
					fmt.Println("")
					continue
				} else if len(gotten) > 3 {
					if gotten[:3] == "-n " {
						gotten = gotten[3:]
						if gotten[1:] == "-" {
							gotten = gotten[1:]
						}
						g, err := strconv.Atoi(gotten)
						if err != nil {
							//fmt.Println(err)
							g = 0
						}
						if g > 255 {
							g -= 255
						}
						bytes[pointer] = byte(g)
					} else if gotten[:3] == "-b " {
						gotten = gotten[3:]
						if gotten[1:] == "-" {
							gotten = gotten[1:]
						}
						g, err := strconv.ParseInt(gotten, 2, 5)
						if err != nil {
							//fmt.Println(err)
							g = 0
						}
						if g > 255 {
							g -= 255
						}
						bytes[pointer] = byte(g)
					} else if gotten[:3] == "-x " {
						gotten = gotten[3:]
						if gotten[1:] == "-" {
							gotten = gotten[1:]
						}
						g, err := strconv.ParseInt(gotten, 16, 5)
						if err != nil {
							//fmt.Println(err)
							g = 0
						}
						if g > 255 {
							g -= 255
						}
						bytes[pointer] = byte(g)
					} else {
						if gotten != "-s" {
							bytes[pointer] = byte(gotten[0])
						}
					}
				} else {
					if gotten != "-s" {
						bytes[pointer] = byte(gotten[0])
					}
				}
				break
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

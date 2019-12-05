package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	tokens := strings.Split(lines[0], ",")
	var mem = []int{}

	for _, i := range tokens {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		mem = append(mem, j)
	}

	input := 1
	currentInstruction := 0

	// defer fmt.Println(mem)

	for {
		mode1 := false
		mode2 := false

		opcode := mem[currentInstruction]

		if mem[currentInstruction] >= 10 {
			token := strconv.Itoa(mem[currentInstruction])
			len := len(token)
			opcode = int(token[len-1] - '0')

			if len > 2 {
				mode1 = token[len-3] == '1'
			}
			if len > 3 {
				mode2 = token[len-4] == '1'
			}
			if len > 4 {
				mode2 = token[len-5] == '1'
			}
		}

		switch opcode {
		case 1:
			var val1, val2 int

			if mode1 {
				val1 = mem[currentInstruction+1]
			} else {
				val1 = mem[mem[currentInstruction+1]]
			}
			if mode2 {
				val2 = mem[currentInstruction+2]
			} else {
				val2 = mem[mem[currentInstruction+2]]
			}
			mem[mem[currentInstruction+3]] = val2 + val1
			currentInstruction += 4
		case 2:
			var val1, val2 int

			if mode1 {
				val1 = mem[currentInstruction+1]
			} else {
				val1 = mem[mem[currentInstruction+1]]
			}
			if mode2 {
				val2 = mem[currentInstruction+2]
			} else {
				val2 = mem[mem[currentInstruction+2]]
			}
			mem[mem[currentInstruction+3]] = val2 * val1
			currentInstruction += 4
		case 3:
			mem[mem[currentInstruction+1]] = input
			currentInstruction += 2
		case 4:
			fmt.Printf("Output: %d\n", mem[mem[currentInstruction+1]])
			currentInstruction += 2
		case 9:
			return
		default:
			currentInstruction++
		}
	}
}

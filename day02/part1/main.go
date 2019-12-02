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

	currentInstruction := 0

	defer fmt.Println(mem)

	for {
		switch mem[currentInstruction] {
		case 1:
			mem[mem[currentInstruction+3]] = mem[mem[currentInstruction+2]] + mem[mem[currentInstruction+1]]
		case 2:
			mem[mem[currentInstruction+3]] = mem[mem[currentInstruction+2]] * mem[mem[currentInstruction+1]]
		default:
			return
		}
		currentInstruction += 4
	}
}

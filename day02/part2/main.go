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
	mem := []int{}

	for _, i := range tokens {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		mem = append(mem, j)
	}

	memcpy := make([]int, len(mem))

	for i := 0; i < 9999; i++ {
		for j := 0; j < 9999; j++ {
			copy(memcpy, mem)

			memcpy[1] = i
			memcpy[2] = j

			currentInstruction := 0

			for mem[currentInstruction] != 99 {
				switch mem[currentInstruction] {
				case 1:
					if memcpy[currentInstruction+3] >= len(memcpy) || memcpy[currentInstruction+2] >= len(memcpy) || memcpy[currentInstruction+1] >= len(memcpy) {
						break
					}
					memcpy[memcpy[currentInstruction+3]] = memcpy[memcpy[currentInstruction+2]] + memcpy[memcpy[currentInstruction+1]]
				case 2:
					if memcpy[currentInstruction+3] >= len(memcpy) || memcpy[currentInstruction+2] >= len(memcpy) || memcpy[currentInstruction+1] >= len(memcpy) {
						break
					}
					memcpy[memcpy[currentInstruction+3]] = memcpy[memcpy[currentInstruction+2]] * memcpy[memcpy[currentInstruction+1]]
				}
				currentInstruction += 4
			}

			if memcpy[0] == 19690720 {
				fmt.Printf("Found: %d %d", i, j)
				return
			}
		}
	}
}

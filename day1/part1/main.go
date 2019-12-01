package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	ans := 0
	for _, line := range lines {
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatalf("Atoi: %s", err)
		}

		ans += (i / 3) - 2
	}

	fmt.Printf("The result is: %d\n", ans)
}

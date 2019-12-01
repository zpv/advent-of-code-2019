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

	cache := make(map[int]int)

	ans := 0
	for _, line := range lines {
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatalf("Atoi: %s", err)
		}

		ans += calcFuel(i, cache)
	}

	fmt.Printf("The result is: %d\n", ans)
}

func calcFuel(mass int, cache map[int]int) int {
	val, ok := cache[mass]

	if mass <= 0 {
		return 0
	}

	if ok {
		return val
	}

	additionalFuel := (mass / 3) - 2

	if additionalFuel < 0 {
		additionalFuel = 0
	}

	totalFuel := additionalFuel + calcFuel(additionalFuel, cache)
	cache[mass] = totalFuel

	return totalFuel
}

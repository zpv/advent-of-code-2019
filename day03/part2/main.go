package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

type tile struct {
	x, y int
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	line1Tokens := strings.Split(lines[0], ",")
	line2Tokens := strings.Split(lines[1], ",")

	panel := make(map[tile]int)

	// Draw 1st line
	curPos := tile{x: 0, y: 0}
	dist := 0

	for _, move := range line1Tokens {
		dir, num := parseToken(move)
		for num > 0 {
			switch dir {
			case 'U':
				curPos.y++
			case 'R':
				curPos.x++
			case 'D':
				curPos.y--
			case 'L':
				curPos.x--
			}
			dist++
			num--
			panel[curPos] = dist
		}
	}

	curPos = tile{x: 0, y: 0}
	dist = 0
	min := math.MaxInt32

	for _, move := range line2Tokens {
		dir, num := parseToken(move)
		for num > 0 {
			switch dir {
			case 'U':
				curPos.y++
			case 'R':
				curPos.x++
			case 'D':
				curPos.y--
			case 'L':
				curPos.x--
			}
			dist++
			num--

			if val, ok := panel[curPos]; ok {
				totalDist := val + dist

				if totalDist < min {
					min = totalDist
				}
			}
		}
	}

	fmt.Println(min)
}

func parseToken(token string) (byte, int) {
	val, err := strconv.Atoi(token[1:])
	if err != nil {
		panic(err)
	}
	return token[0], val
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

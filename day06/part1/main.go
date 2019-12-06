package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type planet struct {
	orbits *planet
	count  int
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var planets = make(map[string]*planet)

	for _, line := range lines {

		if err != nil {
			log.Fatalf("Atoi: %s", err)
		}

		tokens := strings.Split(line, ")")

		// i oop some nasty stuff
		if oribiter, ok := planets[tokens[1]]; ok {
			if oribitee, ok2 := planets[tokens[0]]; ok2 {
				oribiter.orbits = oribitee
			} else {
				oribitee := &planet{count: -1}
				oribiter.orbits = oribitee
				planets[tokens[0]] = oribitee
			}
		} else {
			oribiter := &planet{count: -1}
			if oribitee, ok2 := planets[tokens[0]]; ok2 {
				oribiter.orbits = oribitee
			} else {
				oribitee := &planet{count: -1}
				oribiter.orbits = oribitee
				planets[tokens[0]] = oribitee
			}
			planets[tokens[1]] = oribiter
		}
	}

	sum := 0

	for _, v := range planets {
		countOrbits(v)
		sum += v.count
	}

	fmt.Println(sum)
}

func countOrbits(p *planet) int {
	if p.count >= 0 {
		return p.count
	}

	if p.orbits != nil {
		p.count = countOrbits(p.orbits) + 1
	} else {
		p.count = 0
	}

	return p.count
}

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

	fmt.Println(findMinimumTransfers(planets["YOU"], planets["SAN"]))
}

func findMinimumTransfers(a, b *planet) int {
	visited := make(map[*planet]int)

	cur := a
	dist := 0

	for cur != nil {
		visited[cur] = dist
		dist++
		cur = cur.orbits
	}

	dist = 0
	cur = b

	_, found := visited[cur]
	for ; !found; _, found = visited[cur] {
		dist++
		cur = cur.orbits
	}

	val := visited[cur]
	return dist + val
}

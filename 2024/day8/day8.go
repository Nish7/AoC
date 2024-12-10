package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	res := parseInput(input)
	freqs := GetAllFrequencies(res)
	antinodes := GetAllAntinodes(freqs, res)
	fmt.Println(freqs, antinodes)
}

func GetAllAntinodes(mp map[string][][2]int, ant [][]string) int {
	antinodes := 0

	addAntinode := func(i, j int) bool {
		if i >= len(ant) || j >= len(ant[0]) || i < 0 || j < 0 {
			return false
		}

		if ant[i][j] == "#" {
			return true
		}

		ant[i][j] = "#"
		antinodes++
		return true
	}

	// loop thorugh each keys
	for _, v := range mp {
		for i := 0; i < len(v)-1; i++ {
			p1 := v[i]
			for j := i + 1; j < len(v); j++ {
				p2 := v[j]
				dx, dy := p2[0]-p1[0], p2[1]-p1[1]

				for r := 0; r < len(ant); r++ {
					if !addAntinode(p1[0]-(r*dx), p1[1]-(r*dy)) {
						break
					}
				}
				for r := 0; r < len(ant); r++ {
					if !addAntinode(p2[0]+(r*dx), p2[1]+(r*dy)) {
						break
					}
				}
			}
		}
	}

	// for printing and debugging purposes
	for _, l := range ant {
		fmt.Println(l)
	}

	return antinodes
}

func GetAllFrequencies(mp [][]string) map[string][][2]int {
	res := make(map[string][][2]int)

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[0]); j++ {
			v := mp[i][j]
			if v != "." {
				res[v] = append(res[v], [2]int{i, j})
			}
		}
	}

	return res
}

func parseInput(st string) [][]string {
	trimNewLine := strings.TrimSuffix(st, "\n")
	lines := strings.Split(trimNewLine, "\n")

	var vals [][]string
	for _, l := range lines {
		fields := strings.Split(l, "")
		vals = append(vals, fields)
	}

	return vals
}

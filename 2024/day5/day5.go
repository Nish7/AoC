package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	rules, updates := parseInput(input)
	goodUpdates := checkUpdates(updates, rules)
	midVal := calcMidValues(goodUpdates)
	fmt.Println(midVal)
}

func calcMidValues(goodUpdates [][]int) int {
	mid := 0
	for _, u := range goodUpdates {
		midI := int(math.Floor(float64((len(u) - 1) / 2)))
		mid += u[midI]
	}
	return mid
}

func checkUpdates(updates [][]int, rules [][2]int) (goodUpdates [][]int) {
	// for each updates; go thorugh each rule and find the index of both rules and if they exist, then compare left should be greater
	for _, up := range updates {
		isNoBeuno := false
		for i := 0; i < len(rules); i++ {
			r := rules[i]
			uno := slices.Index(up, r[0])
			secondo := slices.Index(up, r[1])
			if uno != -1 && secondo != -1 {
				if uno > secondo {
					isNoBeuno = true
					up[uno], up[secondo] = up[secondo], up[uno]
					i = 0
				}
			}
		}

		if isNoBeuno {
			goodUpdates = append(goodUpdates, up)
		}
	}

	return
}

func parseInput(st string) (edges [][2]int, orders [][]int) {
	p2idx := 0
	trimNewLine := strings.TrimSuffix(st, "\n")
	lines := strings.Split(trimNewLine, "\n")

	for i, l := range lines {
		if l == "" {
			p2idx = i
			break
		}
		var n1, n2 int
		fmt.Sscanf(l, "%d|%d", &n1, &n2)
		edges = append(edges, [2]int{n1, n2})
	}

	for i := p2idx + 1; i < len(lines); i++ {
		fields := strings.Split(lines[i], ",")

		var flds []int
		for _, f := range fields {
			fi, _ := strconv.Atoi(f)
			flds = append(flds, fi)
		}
		orders = append(orders, flds)
	}
	return
}

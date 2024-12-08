package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	res := parseInput(input)
	r := GetCalibrationResult(res)
	fmt.Println(r)
}

func GetCalibrationResult(res [][]int) int {
	posibleRes := 0

	for _, vals := range res {
		found := false
		r := vals[0]
		operators := GetOperators(len(vals) - 2)
		for _, po := range operators {
			res := vals[1]
			for j, op := range po {
				switch op {
				case 0:
					res += vals[j+2]
				case 1:
					res *= vals[j+2]
				case 2:
					res, _ = strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(vals[j+2]))
				}
			}

			if res == r {
				found = true
				break
			}
		}

		if found {
			posibleRes += r
		}
	}

	return posibleRes
}

func GetOperators(n int) [][]int {
	var res [][]int

	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) >= n {
			res = append(res, path)
			return
		}

		for _, i := range [3]int{0, 1, 2} {
			cp := append([]int{}, path...)
			cp = append(cp, i)
			backtrack(cp)
		}
	}

	backtrack([]int{})
	return res
}

func parseInput(st string) [][]int {
	trimNewLine := strings.TrimSuffix(st, "\n")
	lines := strings.Split(trimNewLine, "\n")

	var vals [][]int
	for _, l := range lines {
		fields := strings.Split(l, " ")
		var line []int
		for _, v := range fields {
			v = strings.TrimSuffix(v, ":")
			a, _ := strconv.Atoi(v)
			line = append(line, a)
		}

		vals = append(vals, line)
	}

	return vals
}

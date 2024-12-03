package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`) // matches any pattern with `mul(` and matches any digits
	matches := re.FindAllStringSubmatch(input, -1)

	total := 0
	for _, ma := range matches {
		v1, _ := strconv.Atoi(ma[1])
		v2, _ := strconv.Atoi(ma[2])
		total += v1 * v2
	}

	fmt.Println(total)
}

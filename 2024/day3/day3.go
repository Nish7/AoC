package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	sanitised := SeperateDoAndDonts(input)
	out := Mul(sanitised)
	fmt.Println(out)
}

func SeperateDoAndDonts(input string) string {
	trimmedNewLine := strings.ReplaceAll(input, "\n", "")
	dosNewLine := regexp.MustCompile(`\bdo\(\)`).ReplaceAllString(trimmedNewLine, "\ndo()")
	dontNewLine := regexp.MustCompile(`don't\(\)`).ReplaceAllString(dosNewLine, "\ndon't()")

	lines := strings.Split(dontNewLine, "\n")
	var result []string
	for _, line := range lines {
		if !strings.Contains(line, "don't()") {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}

func Mul(input string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`) // matches any pattern with `mul(` and matches any digits
	matches := re.FindAllStringSubmatch(input, -1)

	total := 0
	for _, ma := range matches {
		v1, _ := strconv.Atoi(ma[1])
		v2, _ := strconv.Atoi(ma[2])
		total += v1 * v2
	}

	return total
}

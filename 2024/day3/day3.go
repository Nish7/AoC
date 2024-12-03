package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	// input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	re := regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	total := 0
	for _, ma := range matches {
		v1, _ := strconv.Atoi(ma[1])
		v2, _ := strconv.Atoi(ma[2])
		total += v1 * v2
	}

	fmt.Println(total)
}

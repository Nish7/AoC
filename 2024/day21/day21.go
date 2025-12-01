package main

import (
	_ "embed"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

// current pos: [3, 2]
// dest: [0, 0]
func main() {
	input := ParseInput(input)
	out := input.Run()
	return 0
}

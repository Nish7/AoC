package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	input := ParseInput(input)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == "S" {
				fmt.Println("Starting", i, j)
				out := FindPath(i, j, input)
				fmt.Println(out)
				break
			}
		}
	}
}

func FindPath(sx, sy int, board [][]string) int {
	min_v := math.MaxInt64
	seen := make(map[[2]int]bool)

	directions := []struct {
		dx, dy int
		dir    string
	}{
		{1, 0, "S"},  // South
		{-1, 0, "N"}, // North
		{0, 1, "E"},  // East
		{0, -1, "W"}, // West
	}

	queue := [][4]interface{}{{sx, sy, 1001, ""}}
	seen[[2]int{sx, sy}] = true

	for len(queue) > 0 {
		sort.SliceStable(queue, func(i, j int) bool {
			return queue[i][2].(int) < queue[j][2].(int)
		})

		node := queue[0]
		queue = queue[1:]

		fmt.Println(node)

		i, j, p, lastDir := node[0].(int), node[1].(int), node[2].(int), node[3].(string)

		for _, d := range directions {
			points := p + 1

			if lastDir != "" && d.dir != lastDir {
				points += 1000
			}

			ni, nj := i+d.dx, j+d.dy

			if ni >= 0 && nj >= 0 && ni < len(board) && nj < len(board[0]) && board[ni][nj] != "#" && !seen[[2]int{ni, nj}] {
				if board[ni][nj] == "E" {
					if p < min_v {
						min_v = p
					}
					continue
				}
				seen[[2]int{ni, nj}] = true
				queue = append(queue, [4]interface{}{ni, nj, points, d.dir})
			}
		}
	}

	return min_v
}

func MinInt(a, b, c, d int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	if d < min {
		min = d
	}

	return min
}

func ParseInput(in string) [][]string {
	in = strings.TrimSuffix(in, "\n")
	var out [][]string
	fields := strings.Split(in, "\n")
	for _, l := range fields {
		vals := strings.Split(l, "")
		out = append(out, vals)
	}
	return out
}

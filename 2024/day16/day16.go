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
				out, lenOp := FindPath(i, j, input)
				fmt.Println(out, lenOp)
				break
			}
		}
	}
}

var directions = []struct {
	dx, dy int
	dir    string
}{
	{1, 0, "S"},  // South
	{-1, 0, "N"}, // North
	{0, 1, "E"},  // East
	{0, -1, "W"}, // West
}

type VisitedNode struct {
	x   int
	y   int
	dir string
}

type QueueNode struct {
	x, y, score int
	dir         string
	path        map[[2]int]bool
}

func FindPath(sx, sy int, board [][]string) (int, int) {
	min_v := math.MaxInt64
	seen := make(map[VisitedNode]int)
	queue := []QueueNode{{sx, sy, 1000, "", map[[2]int]bool{{sx, sy}: true}}} // we can add all the
	seen[VisitedNode{sx, sy, ""}] = 0
	optimalNodes := make(map[[2]int]bool)

	for len(queue) > 0 {

		sort.SliceStable(queue, func(i, j int) bool { // dijkastra's portion
			return queue[i].score < queue[j].score // ideally should you be using a heap
		})

		node := queue[0]
		queue = queue[1:]
		i, j, p, dir, path := node.x, node.y, node.score, node.dir, node.path

		if board[i][j] == "E" {
			if p <= min_v {
				min_v = p
				for k := range path {
					optimalNodes[k] = true
				}
			}
			continue
		}

		// time.Sleep(0 * time.Millisecond)
		board[i][j] = "O"
		// PrintBoard(board)
		// fmt.Println(node)

		for _, d := range directions {
			points := p + 1
			if dir != "" && d.dir != dir {
				points += 1000
			}

			ni, nj := i+d.dx, j+d.dy

			if ni < 0 ||
				nj < 0 ||
				ni >= len(board) ||
				nj >= len(board[0]) ||
				points > min_v ||
				board[ni][nj] == "#" {
				continue
			}

			if v, ok := seen[VisitedNode{ni, nj, d.dir}]; ok && v < points {
				continue
			}

			newPath := make(map[[2]int]bool)
			for k := range path {
				newPath[k] = true
			}
			newPath[[2]int{ni, nj}] = true
			seen[VisitedNode{ni, nj, d.dir}] = points
			queue = append(queue, QueueNode{ni, nj, points, d.dir, newPath})

		}
	}

	return min_v, len(optimalNodes)
}

func PrintBoard(input [][]string) {
	fmt.Print("\x1b[2J\x1b[H")

	var builder strings.Builder

	for _, row := range input {
		for _, c := range row {
			switch c {
			case "O":
				builder.WriteString("\x1b[31m@\x1b[0m") // Red for '@'
			case "#":
				builder.WriteString("\x1b[32m" + c + "\x1b[0m") // Green for '[' and ']'
			default:
				builder.WriteString(c) // Default character
			}
		}
		builder.WriteByte('\n') // Newline after each row
	}

	fmt.Print(builder.String())
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

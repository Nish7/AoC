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
	board := parseInput(input)
	xmasCount := wordSearch(board, "XMAS")
	fmt.Printf("%d\n", xmasCount)
	o := xMAS(board)
	fmt.Printf("xMAS Count: %d\n", o)
}

func xMAS(board [][]string) int {
	checkMAS := func(i, j int) int {
		if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 {
			return 0
		}

		switch board[i][j] {
		case "S":
			return 1
		case "M":
			return 2
		}
		return 0
	}

	count := 0
	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == "A" {
				if (checkMAS(i+1, j+1)+checkMAS(i-1, j-1) == 3) && (checkMAS(i+1, j-1)+checkMAS(i-1, j+1) == 3) {
					fmt.Println(i, j)
					count++
				}
			}
		}
	}

	return count
}

func wordSearch(board [][]string, word string) int {
	count := 0

	var search func(x, y, idx int, path [][2]int, xDir int, yDir int)
	search = func(x, y, idx int, path [][2]int, xDir int, yDir int) {
		if idx == len(word) {
			count++
			return
		}

		if x >= len(board) || y >= len(board[1]) || x < 0 || y < 0 || board[x][y] != string(word[idx]) || contains(path, [2]int{x, y}) {
			return
		}

		path = append(path, [2]int{x, y})
		search(x+xDir, y+yDir, idx+1, path, xDir, yDir)
	}

	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == string(word[0]) {
				path := [][2]int{}
				search(i, j, 0, path, 1, 0)
				search(i, j, 0, path, 0, 1)
				search(i, j, 0, path, -1, 0)
				search(i, j, 0, path, 0, -1)
				search(i, j, 0, path, 1, -1)
				search(i, j, 0, path, -1, -1)
				search(i, j, 0, path, 1, 1)
				search(i, j, 0, path, -1, 1)
			}
		}
	}

	return count
}

func contains(path [][2]int, coord [2]int) bool {
	for _, p := range path {
		if p == coord {
			return true
		}
	}
	return false
}

func parseInput(st string) [][]string {
	trimNewLine := strings.TrimSuffix(st, "\n")
	lines := strings.Split(trimNewLine, "\n")
	grid := make([][]string, len(lines))
	for i, l := range lines {
		chars := strings.Split(l, "") // [m, m, s, x]
		grid[i] = chars
	}

	return grid
}

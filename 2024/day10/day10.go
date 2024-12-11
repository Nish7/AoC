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
	grid := parseGrid(test)
	// fmt.Println(grid)
	scores := Hike(grid)
	total := Accumulate(scores)
	fmt.Println(scores, total)
}

func Accumulate(scores []int) int {
	sum := 0
	for _, l := range scores {
		sum += l
	}
	return sum
}

func Hike(grid [][]int) []int {
	var res []int
	var dfs func(i, j, n int) int
	dfs = func(i, j, n int) int {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != n {
			return 0
		}

		if n == 9 && grid[i][j] == 9 {
			return 1
		}

		return dfs(i+1, j, n+1) +
			dfs(i-1, j, n+1) +
			dfs(i, j+1, n+1) +
			dfs(i, j-1, n+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				res = append(res, dfs(i, j, 0))
			}
		}
	}

	return res
}

func parseGrid(in string) [][]int {
	in = strings.TrimSuffix(in, "\n")
	fields := strings.Split(in, "\n")

	var grid [][]int
	for _, f := range fields {
		vals := strings.Split(f, "")
		var row []int
		for _, val := range vals {
			v, _ := strconv.Atoi(val)
			row = append(row, v)
		}

		grid = append(grid, row)
	}

	return grid
}

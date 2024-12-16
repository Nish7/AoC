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
	garden := GetGarden(input)
	plots := GetPlots(garden)
	fmt.Println(plots)
}

func GetPlots(garden [][]string) int {
	visited := make(map[[2]int]bool)
	res := 0

	isSame := func(x, y int, z string, grid [][]string) bool {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return false
		}
		return grid[x][y] == z
	}

	getCorners := func(i, j int) int {
		z := garden[i][j]

		NW := isSame(i-1, j-1, z, garden) // Northwest
		W := isSame(i, j-1, z, garden)    // West
		SW := isSame(i+1, j-1, z, garden) // Southwest
		N := isSame(i-1, j, z, garden)    // North
		S := isSame(i+1, j, z, garden)    // South
		NE := isSame(i-1, j+1, z, garden) // Northeast
		E := isSame(i, j+1, z, garden)    // East
		SE := isSame(i+1, j+1, z, garden) // Southeast

		corners := 0

		if N && W && !NW {
			corners++
		}
		if N && E && !NE {
			corners++
		}
		if S && W && !SW {
			corners++
		}
		if S && E && !SE {
			corners++
		}

		if !(N || W) {
			corners++
		}
		if !(N || E) {
			corners++
		}
		if !(S || W) {
			corners++
		}
		if !(S || E) {
			corners++
		}

		return corners
	}

	var dfs func(i, j int, ch string) [3]int
	dfs = func(i, j int, ch string) [3]int {
		if i >= len(garden) || j >= len(garden[0]) || i < 0 || j < 0 || garden[i][j] != ch {
			return [3]int{0, 1, 0}
		}

		if _, ok := visited[[2]int{i, j}]; ok {
			return [3]int{0, 0, 0}
		}

		visited[[2]int{i, j}] = true
		down := dfs(i+1, j, ch)
		up := dfs(i-1, j, ch)
		right := dfs(i, j+1, ch)
		left := dfs(i, j-1, ch)

		area := up[0] + down[0] + right[0] + left[0] + 1
		perimeter := up[1] + down[1] + right[1] + left[1]
		corners := up[2] + down[2] + right[2] + left[2] + getCorners(i, j)

		return [3]int{area, perimeter, corners}
	}

	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[0]); j++ {
			if ok := visited[[2]int{i, j}]; !ok {
				plotVals := dfs(i, j, garden[i][j])
				res += plotVals[0] * plotVals[2]
			}
		}
	}

	return res
}

func GetGarden(in string) [][]string {
	in = strings.TrimSuffix(in, "\n")

	var grid [][]string
	for _, f := range strings.Split(in, "\n") {
		rows := strings.Split(f, "")
		grid = append(grid, rows)
	}

	return grid
}

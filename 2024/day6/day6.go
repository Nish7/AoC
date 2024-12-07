package main

import (
	_ "embed"
	"errors"
	"fmt"
	"slices"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	grid := GetGrid(input)

	position, direction, err := GetStartingPosition(grid)
	if err != nil {
		fmt.Println(err)
	}

	positions, isLoop := MoveGuard(position, direction, grid)
	fmt.Println(len(positions), isLoop)
	loopPositions := PutLoopObstacle(position, direction, positions, grid)
	fmt.Println(loopPositions)

}

func PutLoopObstacle(startPosition [2]int, direction [2]int, positions [][2]int, grid [][]string) int {
	lp := 0
	for _, p := range positions {
		if p[0] == startPosition[0] && p[1] == startPosition[1] {
			continue
		}

		grid[p[0]][p[1]] = "#"
		_, isLoop := MoveGuard(startPosition, direction, grid)

		if isLoop {
			lp++
		}

		grid[p[0]][p[1]] = "."
	}

	return lp
}

func MoveGuard(position [2]int, direction [2]int, grid [][]string) ([][2]int, bool) {
	// apprpach: start with the intial positiong and direction keep switching the directions by 90deg once you find a obstacle and end the while loop when gets out of board
	i, j := position[0], position[1]

	var visited [][2][2]int
	var visitedDir [][2]int
	visited = append(visited, [2][2]int{{i, j}, {direction[0], direction[1]}}) // tracks visited cells with direction
	visitedDir = append(visitedDir, [2]int{i, j})                              // track unique visited cells

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for {
		nexti := i + direction[0]
		nextj := j + direction[1]
		if nexti >= len(grid) || nexti < 0 || nextj >= len(grid[0]) || nextj < 0 {
			break
		}

		if grid[nexti][nextj] == "#" {
			direction = RotateDirection(direction, directions)
			continue
		}

		i = nexti
		j = nextj
		coor := [2][2]int{{i, j}, {direction[0], direction[1]}}
		coorD := [2]int{i, j}

		if !slices.Contains(visitedDir, coorD) {
			visitedDir = append(visitedDir, coorD)
		}

		if !slices.Contains(visited, coor) {
			visited = append(visited, coor)
		} else {
			return visitedDir, true
		}
	}

	return visitedDir, false
}

func RotateDirection(current [2]int, directions [][2]int) [2]int {
	for idx, dir := range directions {
		if dir == current {
			return directions[(idx+1)%len(directions)]
		}
	}
	return current
}

func GetStartingPosition(grid [][]string) ([2]int, [2]int, error) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch grid[i][j] {
			case ">":
				return [2]int{i, j}, [2]int{0, 1}, nil
			case "v":
				return [2]int{i, j}, [2]int{1, 0}, nil
			case "<":
				return [2]int{i, j}, [2]int{0, -1}, nil
			case "^":
				return [2]int{i, j}, [2]int{-1, 0}, nil
			}
		}
	}

	return [2]int{}, [2]int{}, errors.New("No Guard Found")
}

func GetGrid(st string) [][]string {
	trimNewLine := strings.TrimSuffix(st, "\n")
	lines := strings.Split(trimNewLine, "\n")

	var grid [][]string
	for _, l := range lines {
		fields := strings.Split(l, "")
		grid = append(grid, fields)
	}

	return grid
}

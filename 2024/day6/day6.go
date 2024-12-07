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
	grid := GetGrid(test)
	fmt.Println(grid)

	position, direction, err := GetStartingPosition(grid)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(position)
	positions := MoveGuard(position, direction, grid)
	fmt.Println(positions)
}

func MoveGuard(position [2]int, direction [2]int, grid [][]string) int {
	// apprpach: start with the intial positiong and direction keep switching the directions by 90deg once you find a obstacle and end the while loop when gets out of board
	i, j := position[0], position[1]
	var visited [][2]int
	visited = append(visited, [2]int{i, j})
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for {

		fmt.Println(i, j, direction, grid[i][j])
		// check if next coord not outside
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
		coor := [2]int{i, j}
		grid[i][j] = "X"

		if !slices.Contains(visited, coor) {
			visited = append(visited, coor)
		}
	}

	return len(visited)
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

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
	board, directions := ParseInput(input)
	var sx, sy int

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "@" {
				sx, sy = i, j
			}
		}
	}

	fmt.Println("Starting Position: ", sx, sy)

	Move := func(xdir, ydir int) {
		frontX, frontY := sx+xdir, sy+ydir

		if frontX < 0 || frontX >= len(board) || frontY < 0 || frontY >= len(board[0]) {
			return
		}

		inFront := board[frontX][frontY]
		if inFront == "#" {
			return
		}

		if inFront == "." {
			board[sx][sy], board[frontX][frontY] = board[frontX][frontY], board[sx][sy]
			sx += xdir
			sy += ydir
			return
		}

		if inFront == "O" {
			var boxes [][2]int
			currentX, currentY := frontX, frontY

			for { // keep looping in the same dir until u find the boxes
				if currentX < 0 || currentX >= len(board) || currentY < 0 || currentY >= len(board[0]) {
					break
				}

				cell := board[currentX][currentY]
				if cell == "O" {
					boxes = append(boxes, [2]int{currentX, currentY})
					currentX += xdir
					currentY += ydir
				} else {
					break
				}
			}

			if currentX < 0 || currentX >= len(board) || currentY < 0 || currentY >= len(board[0]) {
				return
			}

			// check if the next cell is .
			nextCell := board[currentX][currentY]
			if nextCell != "." {
				return
			}

			// move last boxes one by one
			for i := len(boxes) - 1; i >= 0; i-- {
				box := boxes[i]
				newX, newY := box[0]+xdir, box[1]+ydir
				board[newX][newY] = "O"
				board[box[0]][box[1]] = "."
			}

			board[sx][sy], board[frontX][frontY] = board[frontX][frontY], board[sx][sy]
			sx += xdir
			sy += ydir
			return
		}
	}

	for _, d := range directions {
		// fmt.Println("----- Taking this Direction: ", d)
		switch d {
		case "<":
			Move(0, -1)
		case "^":
			Move(-1, 0)
		case "v":
			Move(1, 0)
		case ">":
			Move(0, 1)
		}
	}

	// PrintBoard(board)
	val := TotalGPSCoord(board)
	fmt.Println(val)
}

func PrintBoard(board [][]string) {
	for _, row := range board {
		joinedRow := strings.Join(row, " ")
		fmt.Println(joinedRow)
	}
}

func TotalGPSCoord(board [][]string) int {
	gps := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "O" {
				gps += 100*i + j
			}
		}
	}

	return gps
}

func ParseInput(in string) (board [][]string, directions []string) {
	parts := strings.Split(in, "\n\n")

	p1 := parts[0]
	p2 := strings.ReplaceAll(parts[1], "\n", "")
	directions = strings.Split(p2, "")

	fields := strings.Split(p1, "\n")
	var rows [][]string
	for _, f := range fields {
		v := strings.Split(f, "")
		var row []string
		for _, inp := range v {
			var ch string
			switch inp {

			}
			row = append(row, ch)
		}
		row = append(row, v)
	}

	return row, directions
}

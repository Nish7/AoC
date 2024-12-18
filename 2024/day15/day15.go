package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
	// "time"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	board, directions := ParseInput(test)
	var sx, sy int

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "@" {
				sx, sy = i, j
			}
		}
	}

	fmt.Println("Starting Position: ", sx, sy)
	PrintBoard(board)

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

		// Handle box pushing
		type Box struct {
			x1 int
			y1 int
			x2 int
			y2 int
		}

		var leadingChar string
		switch {
		case xdir == 0 && ydir == -1: // Left
			leadingChar = "]"
		case xdir == 0 && ydir == 1: // Right
			leadingChar = "["
		}

		if inFront == leadingChar && ydir != 0 {
			var boxes []Box
			closingX, closingY := frontX, frontY+ydir
			boxes = append(boxes, Box{frontX, frontY, closingX, closingY})

			currentXa, currentYa := closingX, closingY+ydir
			currentXb, currentYb := currentXa, currentYa+ydir
			// FIX: currently it is moving in two pairs. that doesnt work as we should be only looking into moving into next one
			for {
				if currentXa > len(board) || currentYa > len(board[0]) || currentXa < 0 || currentYa < 0 {
					break
				}

				if currentXb > len(board) || currentYb > len(board[0]) || currentXb < 0 || currentYb < 0 {
					break
				}

				if board[currentXa][currentYa] != leadingChar && board[currentXb][currentYb] != leadingCharFlipped(leadingChar) {
					break
				}

				boxes = append(boxes, Box{currentXa, currentYa, currentXb, currentYb})

				currentYa = currentYa + (2 * ydir)
				currentYb = currentYb + (2 * ydir)
			}

			fmt.Println(board[currentXa][currentYa], board[currentXb][currentYb])
			fmt.Println(boxes)

			if board[currentXb][currentYb] != "." && board[currentXa][currentYa] != "." {
				return
			}

			// move the boxes
			for i := len(boxes) - 1; i >= 0; i-- {
				box := boxes[i]
				newXa, newYa := box.x1+xdir, box.y1+ydir
				newXb, newYb := box.x2+xdir, box.y2+ydir
				board[newXb][newYb] = leadingCharFlipped(leadingChar)
				board[box.x2][box.y2] = "."
				board[newXa][newYa] = leadingChar
				board[box.x1][box.y1] = "."
			}

			board[sx][sy], board[frontX][frontY] = board[frontX][frontY], board[sx][sy]
			sx += xdir
			sy += ydir
		}

		if inFront == "[" || inFront == "]" && xdir != 0 {
			var leftRigth int
			if inFront == "[" {
				leftRigth = 1
			} else {
				leftRigth = -1
			}
			var boxes []Box
			closingX, closingY := frontX, frontY+leftRigth
			boxes = append(boxes, Box{frontX, frontY, closingX, closingY})

			currentXa, currentYa := frontX+xdir, frontY
			currentXb, currentYb := currentXa, closingY
			for {
				if currentXa > len(board) || currentYa > len(board[0]) || currentXa < 0 || currentYa < 0 {
					break
				}

				if currentXb > len(board) || currentYb > len(board[0]) || currentXb < 0 || currentYb < 0 {
					break
				}

				if board[currentXa][currentYa] != inFront && board[currentXb][currentYb] != leadingCharFlipped(inFront) {
					break
				}

				boxes = append(boxes, Box{currentXa, currentYa, currentXb, currentYb})

				currentXa = currentXa + xdir
				currentXb = currentXa
			}

			fmt.Println(board[currentXa][currentYa], board[currentXb][currentYb])
			fmt.Println(boxes)

			if board[currentXb][currentYb] != "." || board[currentXa][currentYa] != "." {
				return
			}

			// move the boxes
			for i := len(boxes) - 1; i >= 0; i-- {
				box := boxes[i]
				newXa, newYa := box.x1+xdir, box.y1+ydir
				newXb, newYb := box.x2+xdir, box.y2+ydir
				board[newXb][newYb] = leadingCharFlipped(inFront)
				board[box.x2][box.y2] = "."
				board[newXa][newYa] = inFront
				board[box.x1][box.y1] = "."
			}

			board[sx][sy], board[frontX][frontY] = board[frontX][frontY], board[sx][sy]
			sx += xdir
			sy += ydir
		}

	}

	for i, d := range directions {
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

		fmt.Print("\033[H\033[2J")
		PrintBoard(board)
		fmt.Println("----- Taking this Direction: ", d)
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}

	val := TotalGPSCoord(board)
	fmt.Println(val)
}

func leadingCharFlipped(char string) string {
	if char == "[" {
		return "]"
	}
	return "["
}

func PrintBoard(board [][]string) {
	for _, row := range board {
		joinedRow := strings.Join(row, "")
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

	p2 := strings.ReplaceAll(parts[1], "\n", "")
	directions = strings.Split(p2, "")

	fields := strings.Split(parts[0], "\n")
	var rows [][]string
	for _, f := range fields {
		v := strings.Split(f, "")
		var row []string
		for _, inp := range v {
			switch inp {
			case "#":
				row = append(row, "#")
				row = append(row, "#")
			case "O":
				row = append(row, "[")
				row = append(row, "]")
			case ".":
				row = append(row, ".")
				row = append(row, ".")
			case "@":
				row = append(row, "@")
				row = append(row, ".")
			}
		}
		rows = append(rows, row)
	}

	return rows, directions
}

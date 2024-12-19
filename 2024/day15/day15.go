package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	// PrintBoard(board)

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

			// fmt.Println(board[currentXa][currentYa], board[currentXb][currentYb])
			// fmt.Println(boxes)

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

		if (inFront == "[" || inFront == "]") && xdir != 0 {
			var boxes []Box
			closingX, closingY := frontX, frontY+getLeftRightVal(board[frontX][frontY])
			boxes = append(boxes, Box{frontX, frontY, closingX, closingY})

			var queue []Box
			queue = append(queue, Box{frontX, frontY, closingX, closingY})
			// c := 0

			for len(queue) > 0 {
				// if len(queue) > 10 || c > 10 {
				// 	break
				// }

				fmt.Println(queue)

				b := queue[0]
				queue = queue[1:]
				// c++

				fmt.Println(b)
				if b.x1+xdir > len(board) || b.x1+xdir < 0 || board[b.x1+xdir][b.y1] == "#" {
					return
				}

				if b.x2+xdir > len(board) || b.x2+xdir < 0 || board[b.x2+xdir][b.y2] == "#" {
					return
				}

				// if the one in front of the opening bracket is same. it means it this box can push one only
				if board[b.x1][b.y1] == board[b.x1+xdir][b.y1] {
					// fmt.Println("bracket same")
					// PrintBoard(board)

					boxes = append(boxes, Box{b.x1 + xdir, b.y1, b.x2 + xdir, b.y2})
					queue = append(queue, Box{b.x1 + xdir, b.y1, b.x2 + xdir, b.y2})
					continue
				}

				// add the box in front of opening bracket if any
				if board[b.x1+xdir][b.y1] == "]" || board[b.x1+xdir][b.y1] == "[" {
					// fmt.Println("something above x1")
					if !slices.Contains(queue, Box{b.x1 + xdir, b.y1 + getLeftRightVal(board[b.x1+xdir][b.y1]), b.x2 + xdir, b.y1}) {
						fmt.Println("MAYBE HERER")
						boxes = append(boxes, Box{b.x1 + xdir, b.y1, b.x2 + xdir, b.y1 + getLeftRightVal(board[b.x1+xdir][b.y1])})
						queue = append(queue, Box{b.x1 + xdir, b.y1, b.x2 + xdir, b.y1 + getLeftRightVal(board[b.x1+xdir][b.y1])})
					}
				}

				// add the obx in front of closing bracker if any
				if board[b.x2+xdir][b.y2] == "]" || board[b.x2+xdir][b.y2] == "[" {
					// fmt.Println("something above x2")
					if !slices.Contains(queue, Box{b.x2 + xdir, b.y2 + getLeftRightVal(board[b.x2+xdir][b.y2]), b.x2 + xdir, b.y2}) {
						fmt.Println("HERER")
						boxes = append(boxes, Box{b.x2 + xdir, b.y2, b.x2 + xdir, b.y2 + getLeftRightVal(board[b.x2+xdir][b.y2])})
						queue = append(queue, Box{b.x2 + xdir, b.y2, b.x2 + xdir, b.y2 + getLeftRightVal(board[b.x2+xdir][b.y2])})
					}
				}

			}

			fmt.Println(boxes)

			// move the boxes
			for i := len(boxes) - 1; i >= 0; i-- {
				box := boxes[i]
				newXa, newYa := box.x1+xdir, box.y1+ydir
				newXb, newYb := box.x2+xdir, box.y2+ydir
				board[newXb][newYb] = board[box.x2][box.y2]
				board[newXa][newYa] = board[box.x1][box.y1]
				board[box.x1][box.y1] = "."
				board[box.x2][box.y2] = "."
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

		// fmt.Print("\033[H\033[2J")
		// PrintBoard(board)
		fmt.Println("Instruction :", i)
		// time.Sleep(100 * time.Millisecond)
	}

	val := TotalGPSCoord(board)
	fmt.Println(val)
	PrintBoard(board)
}

func leadingCharFlipped(char string) string {
	if char == "[" {
		return "]"
	}
	return "["
}

func getLeftRightVal(inFront string) int {
	var leftRigth int
	if inFront == "[" {
		leftRigth = 1
	} else {
		leftRigth = -1
	}

	return leftRigth
}

func PrintBoard(warehouse [][]string) {
	// Clear the screen and reset the cursor
	fmt.Print("\x1b[2J\x1b[H")

	// Use a strings.Builder for efficient concatenation
	var builder strings.Builder

	for _, row := range warehouse {
		for _, c := range row {
			switch c {
			case "@":
				builder.WriteString("\x1b[31m@\x1b[0m") // Red for '@'
			case "[", "]":
				builder.WriteString("\x1b[32m" + c + "\x1b[0m") // Green for '[' and ']'
			default:
				builder.WriteString(c) // Default character
			}
		}
		builder.WriteByte('\n') // Newline after each row
	}

	// Print the entire warehouse at once
	fmt.Print(builder.String())
}

func TotalGPSCoord(board [][]string) int {
	gps := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "[" {
				gps += 100*i + (j)
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

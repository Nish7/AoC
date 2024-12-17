package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	in := parseInput(input)
	w := 101
	h := 103
	midw := int(math.Floor(float64(w / 2)))
	midh := int(math.Floor(float64(h / 2)))

	for i := 0; i < 100; i++ {
		q1 := 0
		q2 := 0
		q3 := 0
		q4 := 0
		for _, r := range in {
			x, y := GetEndPosition(r, w, h, i)
			switch {
			case x < midw && y < midh:
				q1++
			case x > midw && y < midh:
				q2++
			case x < midw && y > midh:
				q3++
			case x > midw && y > midh:
				q4++
			}

			print(q1, q2, q3, q4)
		}
	}

	ans := q1 * q2 * q3 * q4
	fmt.Println(ans)
}

func GetEndPosition(r [2][2]int, w, h int, second int) (int, int) {
	x, y, vx, vy := r[0][0], r[0][1], r[1][0], r[1][1]
	fmt.Println(x, y, vx, vy)
	endx := ((x + vx*second) % w)
	if endx < 0 {
		endx += w
	}

	endy := ((y + vy*second) % h)
	if endy < 0 {
		endy += h
	}
	return endx, endy
}

func parseInput(in string) [][2][2]int {
	in = strings.TrimSuffix(in, "\n")
	vals := strings.Split(in, "\n")
	var inputs [][2][2]int

	for _, l := range vals {
		var x, y, vx, vy int
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		inputs = append(inputs, [2][2]int{{x, y}, {vx, vy}})
	}

	return inputs
}

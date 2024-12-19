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
	totalTime := 10000 // Total simulation time in seconds

	for i := 0; i <= totalTime; i++ {

		var points [][2]float64

		for _, r := range in {
			x, y := GetEndPosition(r, w, h, i)
			points = append(points, [2]float64{float64(x), float64(y)})
		}

		var sumX, sumY float64
		for _, p := range points {

			sumX += p[0]
			sumY += p[1]
		}
		centroidX := sumX / float64(len(points))
		centroidY := sumY / float64(len(points))

		var sumSqDist float64
		for _, p := range points {
			dx := p[0] - centroidX
			dy := p[1] - centroidY
			sumSqDist += dx*dx + dy*dy
		}

		var stdDev float64
		stdDev = math.Sqrt(sumSqDist / float64(len(points)))

		if stdDev >= 0 && stdDev <= 30 { // filter all below 30 for clustered positions
			fmt.Println(i, stdDev)
		}

	}
}

func GetEndPosition(r [2][2]int, w, h int, second int) (int, int) {
	x, y, vx, vy := r[0][0], r[0][1], r[1][0], r[1][1]
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

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

type Point struct {
	X float64
	Y float64
}

func main() {
	in := parseInput(input)
	w := 101
	h := 103
	totalTime := 10000 // Total simulation time in seconds

	for i := 0; i <= totalTime; i++ {
		points := []Point{}

		for _, r := range in {
			x, y := GetEndPosition(r, w, h, i)
			points = append(points, Point{X: float64(x), Y: float64(y)})
		}

		var sumDist float64
		count := 0
		for m := 0; m < len(points); m++ {
			for n := m + 1; n < len(points); n++ {
				dist := math.Hypot(points[m].X-points[n].X, points[m].Y-points[n].Y)
				sumDist += dist
				count++
			}
		}

		var sumX, sumY float64
		for _, p := range points {
			sumX += p.X
			sumY += p.Y
		}
		centroidX := sumX / float64(len(points))
		centroidY := sumY / float64(len(points))

		var sumSqDist float64
		for _, p := range points {
			dx := p.X - centroidX
			dy := p.Y - centroidY
			sumSqDist += dx*dx + dy*dy
		}
		var stdDev float64
		if len(points) > 0 {
			stdDev = math.Sqrt(sumSqDist / float64(len(points)))
		}

		if stdDev >= 0 && stdDev <= 30 {
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

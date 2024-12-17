package main

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

type Machine struct {
	button_a [2]float64 // x and y
	button_b [2]float64
	price    [2]float64
}

func main() {
	machines := GetMachines(test)
	clicks := GetClicks(machines)
	fmt.Printf("%.0f", clicks)
}

func GetClicks(machines []Machine) float64 {
	var totalPoints float64
	for _, m := range machines {

		a, b, error := GetButtons(
			m.button_a[0],
			m.button_a[1],
			m.button_b[0],
			m.button_b[1],
			m.price[0],
			m.price[1],
		)

		if error != nil {
			continue
		}

		points := a*3 + b*1
		totalPoints += points
	}

	return totalPoints
}

func GetButtons(aX, bX, aY, bY, prizeA, prizeB float64) (x, y float64, err error) {
	// Cramers Equation
	// https://math.libretexts.org/Bookshelves/Precalculus/Precalculus_(Stitz-Zeager)/08%3A_Systems_of_Equations_and_Matrices/8.05%3A_Determinants_and_Cramers_Rule
	det := aX*bY - bX*aY

	if det == 0 {
		return 0, 0, errors.New("no solution")
	}

	detX := (prizeA * bY) - (prizeB * aY)
	detY := (aX * prizeB) - (bX * prizeA)

	x = detX / det
	y = detY / det

	if x != math.Floor(x) || y != math.Floor(y) {
		return 0, 0, errors.New("no Solution")
	}

	return x, y, nil
}

func GetMachines(in string) []Machine {
	in = strings.TrimSuffix(in, "\n")
	var machines []Machine

	for _, m := range strings.Split(in, "\n\n") {
		var aX, aY, bX, bY, priceX, priceY float64
		f := strings.Split(m, "\n")

		fmt.Sscanf(f[0], "Button A: X+%f, Y+%f", &aX, &aY)
		fmt.Sscanf(f[1], "Button B: X+%f, Y+%f", &bX, &bY)
		fmt.Sscanf(f[2], "Prize: X=%f, Y=%f", &priceX, &priceY)
		priceX += 10000000000000
		priceY += 10000000000000

		machines = append(machines, Machine{
			button_a: [2]float64{aX, aY},
			button_b: [2]float64{bX, bY},
			price:    [2]float64{priceX, priceY},
		})

	}

	return machines

}

package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	p1 := GetAllStones(input)
	fmt.Println(p1)
}

func GetAllStones(inp string) int {
	total := 0
	memo := make(map[[2]int]int)
	s := []int{}

	for _, l := range strings.Split(inp, " ") {
		var Blink func(n int, l int) int
		Blink = func(n int, l int) int { // blink(125) -> 7
			if l == 25 {
				memo[[2]int{n, l}] = 1
				s = append(s, n)
				return 1
			}

			if v, ok := memo[[2]int{n, l}]; ok {
				return v
			}

			if n == 0 {
				val := Blink(1, l+1)
				memo[[2]int{n, l}] = val
				return val
			}

			strn := strconv.Itoa(n)
			if len(strn)%2 == 0 {
				mid := int(math.Floor(float64(len(strn) / 2)))
				p1n, _ := strconv.Atoi(strn[:mid])
				p2n, _ := strconv.Atoi(strn[mid:])
				val := Blink(p1n, l+1) + Blink(p2n, l+1)
				memo[[2]int{n, l}] = val
				return val
			}

			val := Blink(n*2024, l+1)
			memo[[2]int{n, l}] = val
			return val
		}

		l, _ := strconv.Atoi(l)
		stones := Blink(l, 0)
		total += stones
		fmt.Println(stones)
	}

	fmt.Println(s)
	fmt.Println(memo)
	return total
}

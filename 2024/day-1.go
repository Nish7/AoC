package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./puzzle-1-input.txt")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	distance := puzzle1(string(input))
	fmt.Println(distance)

	similarity_score := puzzle2(string(input))
	fmt.Println(similarity_score)
}

func puzzle1(input string) int {
	// get the two lists
	lines := strings.Split(input, "\n")

	var lista []int
	var listb []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		words := strings.Split(line, "   ")
		l1, _ := strconv.Atoi(words[0])
		l2, _ := strconv.Atoi(words[1])
		lista = append(lista, l1)
		listb = append(listb, l2)
	}

	// sort the two lists in ascending order
	sort.Ints(lista)
	sort.Ints(listb)

	// calcualte the difference between them and calucalte a total sum
	distance := 0
	for i, el := range lista {
		distance += int(math.Abs(float64(el - listb[i])))
	}

	return distance
}

func puzzle2(input string) int {
	// get the two lists
	lines := strings.Split(input, "\n")

	var lista []int
	var listb []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		words := strings.Split(line, "   ")
		l1, _ := strconv.Atoi(words[0])
		l2, _ := strconv.Atoi(words[1])
		lista = append(lista, l1)
		listb = append(listb, l2)
	}

	// create a counter map A
	countera := make(map[int]int)
	for _, v := range lista {
		countera[v]++
	}

	counterb := make(map[int]int)
	for _, v := range listb {
		counterb[v]++
	}

	// go through each vals in lista and get the freq of list_b val
	total := 0
	for _, k := range lista {
		v, ok := counterb[k]
		if !ok {
			continue
		}
		total += k * v
	}

	return total

}

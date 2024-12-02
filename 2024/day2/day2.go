package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Reports [][]int

//go:embed input.txt
var input string

func main() {
	reports := parseInput(input)
	fmt.Printf("reports: %v\n", reports)

	safeReports := SafeReports(reports)
	fmt.Printf("Safe Reports Len: %d\n", safeReports)

}

func SafeReports(reports Reports) int {
	r := 0
	for _, report := range reports {
		if IsReportSafe(report) || IsLessSafer(report) {
			fmt.Printf("This is safe : %v\n", report)
			r++
		}
	}

	return r
}

func IsReportSafe(report []int) bool {
	return checkIncrease(report) || checkDecrease(report)
}

func IsLessSafer(report []int) bool {
	fmt.Println(report)
	for i := 0; i < len(report); i++ {
		modified := make([]int, len(report))
		copy(modified, report)
		modified = append(modified[:i], modified[i+1:]...)
		fmt.Printf("checkin this report %v\n", modified)
		if IsReportSafe(modified) {
			return true
		}
	}
	return false
}

func checkIncrease(report []int) bool {
	for r := 1; r < len(report); r++ {
		diff := report[r] - report[r-1]
		if diff != 1 && diff != 2 && diff != 3 {
			return false
		}

	}
	return true
}

func checkDecrease(report []int) bool {
	for r := 1; r < len(report); r++ {
		diff := report[r] - report[r-1]
		if diff != -1 && diff != -2 && diff != -3 {
			return false
		}
	}
	return true
}

func parseInput(input string) Reports {
	lines := strings.Split(input, "\n")
	var reports Reports

	for _, line := range lines {
		if line == "" {
			continue
		}

		values := strings.Fields(line)
		report := make([]int, len(values))

		for i, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error: Parsing the string to int")
			}
			report[i] = num
		}

		reports = append(reports, report)
	}

	return reports
}

package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func main() {
	diskMap := ProcessDiskMap(input)
	compressedMap := CompressDisk(diskMap)
	checksum := FileSystemChecksum(compressedMap)
	fmt.Println(checksum)
}

func FileSystemChecksum(diskMap []string) int {
	checksum := 0
	for i, l := range diskMap {
		if l == "." {
			continue
		}
		v, _ := strconv.Atoi(l)
		checksum += i * v
	}
	return checksum
}

func CompressDisk(diskMap []string) []string {
	right := len(diskMap) - 1

	for right >= 0 {
		// Find the last non-empty block
		for right >= 0 && diskMap[right] == "." {
			right--
		}

		// Calculate the file length
		fileLen := 0
		for right-fileLen >= 0 && diskMap[right-fileLen] == diskMap[right] {
			fileLen++
		}

		left := 0
		for left < right {
			// loop until u find the empty
			for left < right && diskMap[left] != "." {
				left++
			}

			// Count contiguous empty space
			emptySpace := 0
			for left+emptySpace < right && diskMap[left+emptySpace] == "." {
				emptySpace++
			}

			// Check if empty space is sufficient
			if emptySpace >= fileLen {
				// swapping block
				for i := 0; i < fileLen; i++ {
					diskMap[left+i], diskMap[right-i] = diskMap[right-i], "."
				}
				break
			} else {
				left += emptySpace
			}
		}

		// move the right pointer to the next block
		right -= fileLen
	}

	return diskMap
}

func ProcessDiskMap(in string) []string {
	sum := 0
	in = strings.TrimSuffix(in, "\n")
	fields := strings.Split(in, "")
	var diskMapI []string

	for _, f := range fields {
		v, _ := strconv.Atoi(f)
		sum += v
		diskMapI = append(diskMapI, f)
	}

	var diskMap = make([]string, sum)
	nextPointer := 0
	block := 0
	for i, l := range diskMapI {
		l, _ := strconv.Atoi(l)

		if i%2 == 0 {
			// Even indices: place file blocks
			for c := 0; c < l && nextPointer < len(diskMap); c++ {
				diskMap[nextPointer] = strconv.Itoa(block)
				nextPointer++
			}
			block++
		} else {
			// Odd indices: place free spaces
			for c := 0; c < l && nextPointer < len(diskMap); c++ {
				diskMap[nextPointer] = "."
				nextPointer++
			}
		}

	}

	return diskMap
}

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
		v, _ := strconv.Atoi(l)
		checksum += i * v
	}
	return checksum
}

func CompressDisk(diskMap []string) []string {
	left := 0
	right := len(diskMap) - 1

	for {
		if !(left < right) {
			break
		}

		// find the index to start writing into
		for {
			if diskMap[left] == "." || left >= right {
				break
			}
			left++
		}

		// find the index to start reading into
		for {
			if diskMap[right] != "." || right <= left {
				break
			}
			right--
		}

		// keep swapping until
		for {
			if !(left < right && diskMap[left] == "." && diskMap[right] != ".") {
				break
			}

			diskMap[left], diskMap[right] = diskMap[right], diskMap[left]
			left++
			right--
		}
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

		// for even numbers loop the block number until l times
		if i%2 == 0 {
			counter := 0
			for {
				if nextPointer >= len(diskMap) || counter == l {
					break
				}
				diskMap[nextPointer] = strconv.Itoa(block)
				nextPointer += 1
				counter += 1
			}

			block++
		} else {
			// for odd numbers loop the empty space for l times
			counter := 0
			for {
				if nextPointer >= len(diskMap) || counter == l {
					break
				}
				diskMap[nextPointer] = "."
				nextPointer += 1
				counter += 1
			}
		}

	}

	return diskMap
}

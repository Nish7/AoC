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
	fmt.Println(compressedMap, checksum)
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
	right := len(diskMap) - 1

	for right >= 0 {
		// Find the last non-empty block
		for right >= 0 && diskMap[right] == "." {
			right--
		}

		if right < 0 {
			break
		}

		// Calculate the file length
		fileLen := 0
		for fileLen = 0; right-fileLen >= 0 && diskMap[right-fileLen] == diskMap[right]; fileLen++ {
		}

		fmt.Println("Looking a space for ", diskMap[right], fileLen, right)

		left := 0
		for left < right {
			// loop until u find the empty
			for left < right && diskMap[left] != "." {
				left++
			}

			if left >= right {
				break
			}

			// Count contiguous empty space
			emptySpace := 0
			for emptySpace = 0; left+emptySpace < right && diskMap[left+emptySpace] == "."; emptySpace++ {
			}

			if emptySpace < fileLen {
				left += fileLen + 1
				continue
			}

			// Check if empty space is sufficient
			if emptySpace >= fileLen {
				fmt.Println("Swapping block", diskMap[right], "to position", left)
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

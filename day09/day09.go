package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func makeFragmentMap(diskMap string) ([]string, int) {
	checksum := []string{}
	id := 0
	for i := 0; i < len(diskMap); i++ {
		num, _ := strconv.Atoi(string(diskMap[i]))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				checksum = append(checksum, strconv.Itoa(id))
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				checksum = append(checksum, ".")
			}
		}
	}
	return checksum, id - 1
}

func sumChecksum(checksum []string) int {
	sum := 0
	for i := 0; i < len(checksum); i++ {
		if checksum[i] != "." {
			num, _ := strconv.Atoi(checksum[i])
			sum += (num * i)
		}
	}
	return sum
}

func partOne(diskMap string) int {
	checksum, _ := makeFragmentMap(diskMap)

	idxLast := len(checksum) - 1
	for i := 0; i < len(checksum); i++ {
		if idxLast < i {
			break
		}
		if checksum[i] == "." {
			for checksum[idxLast] == "." && idxLast > i {
				idxLast--
			}
			checksum[i] = checksum[idxLast]
			checksum[idxLast] = "."
			idxLast--
		}
	}

	return sumChecksum(checksum)
}

func getIdOcurrence(checksum []string, id string) (int, int) {
	idx := -1
	count := 0

	for i := len(checksum) - 1; i >= 0; i-- {
		if id == checksum[i] {
			count++
			idx = i
		}
	}

	return idx, count
}

func partTwo(diskMap string) int {
	checksum, maxId := makeFragmentMap(diskMap)

	for maxId > 0 {
		index, fileCount := getIdOcurrence(checksum, strconv.Itoa(maxId))

		spaceIdx := -1
		for i := 0; i < len(checksum)-1; i++ {
			if checksum[i] == "." {
				spaceCount := 0
				for i < len(checksum)-1 && checksum[i] == "." {
					spaceCount++
					i++
				}
				if spaceCount >= fileCount {
					spaceIdx = i - spaceCount
					break
				}
			}
		}

		if spaceIdx > -1 && spaceIdx < index {
			for i := spaceIdx; i < spaceIdx+fileCount; i++ {
				checksum[i] = strconv.Itoa(maxId)
			}

			for i := index; i < index+fileCount; i++ {
				checksum[i] = "."
			}
		}

		maxId--
	}

	return sumChecksum(checksum)
}

func main() {
	inputFile, _ := os.Open("day09/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var diskMap string

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		diskMap = scanner.Text()
	}

	resultPartOne := partOne(diskMap)
	fmt.Println("Part one:", resultPartOne)

	resultPartTwo := partTwo(diskMap)
	fmt.Println("Part two:", resultPartTwo)
}

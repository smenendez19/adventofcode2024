package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coordinates struct {
	x int
	y int
}

func findStartPositions(tophographicMap []string) []Coordinates {
	startPositions := []Coordinates{}

	for i := 0; i < len(tophographicMap); i++ {
		for j := 0; j < len(tophographicMap[i]); j++ {
			if string(tophographicMap[i][j]) == "0" {
				startPositions = append(startPositions, Coordinates{i, j})
			}
		}
	}

	return startPositions
}

func findTrailHead(tophographicMap []string, position Coordinates, height int) []Coordinates {
	finishPosition := []Coordinates{}
	if height == 9 {
		return append(finishPosition, position)
	}
	if position.x+1 < len(tophographicMap) && string(tophographicMap[position.x+1][position.y]) == strconv.Itoa(height+1) {
		finishPosition = append(finishPosition, findTrailHead(tophographicMap, Coordinates{position.x + 1, position.y}, height+1)...)
	}
	if position.x-1 > -1 && string(tophographicMap[position.x-1][position.y]) == strconv.Itoa(height+1) {
		finishPosition = append(finishPosition, findTrailHead(tophographicMap, Coordinates{position.x - 1, position.y}, height+1)...)
	}
	if position.y+1 < len(tophographicMap) && string(tophographicMap[position.x][position.y+1]) == strconv.Itoa(height+1) {
		finishPosition = append(finishPosition, findTrailHead(tophographicMap, Coordinates{position.x, position.y + 1}, height+1)...)
	}
	if position.y-1 > -1 && string(tophographicMap[position.x][position.y-1]) == strconv.Itoa(height+1) {
		finishPosition = append(finishPosition, findTrailHead(tophographicMap, Coordinates{position.x, position.y - 1}, height+1)...)
	}
	return finishPosition
}

func setUniquePositions(positions []Coordinates) []Coordinates {
	uniquePositions := map[Coordinates]bool{}
	for i := 0; i < len(positions); i++ {
		uniquePositions[positions[i]] = true
	}
	keys := make([]Coordinates, 0, len(uniquePositions))
	for k := range uniquePositions {
		keys = append(keys, k)
	}
	return keys
}

func partOne(tophographicMap []string) int {
	sum := 0

	startPositions := findStartPositions(tophographicMap)

	for _, startPosition := range startPositions {
		sum += len(setUniquePositions(findTrailHead(tophographicMap, startPosition, 0)))
	}

	return sum
}

func partTwo(tophographicMap []string) int {
	sum := 0

	startPositions := findStartPositions(tophographicMap)

	for _, startPosition := range startPositions {
		sum += len(findTrailHead(tophographicMap, startPosition, 0))
	}

	return sum
}

func main() {
	inputFile, _ := os.Open("day10/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	tophographicMap := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		tophographicMap = append(tophographicMap, scanner.Text())
	}

	resultPartOne := partOne(tophographicMap)
	fmt.Println("Part one:", resultPartOne)

	resultPartTwo := partTwo(tophographicMap)
	fmt.Println("Part two:", resultPartTwo)
}

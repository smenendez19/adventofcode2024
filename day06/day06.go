package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func findGuard(lab []string) (int, int) {
	for i := 0; i < len(lab); i++ {
		for j := 0; j < len(lab[i]); j++ {
			if string(lab[i][j]) == "^" || string(lab[i][j]) == "v" || string(lab[i][j]) == ">" || string(lab[i][j]) == "<" {
				return i, j
			}
		}
	}
	return -1, -1
}

func drawChar(lab []string, x int, y int, c rune) {
	out := []rune(lab[x])
	out[y] = c
	lab[x] = string(out)
}

func drawLab(lab []string) {
	for i := 0; i < len(lab); i++ {
		fmt.Println(lab[i])
	}
}

func rotateGuard(view rune) rune {
	switch view {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	default:
		return view
	}
}

func countX(lab []string) int {
	count := 1

	for _, row := range lab {
		count += strings.Count(row, "X")
	}
	return count
}

func markPositions(lab []string) {
	guardX, guardY := findGuard(lab)
	view := '^'
	for guardX > 0 && guardY > 0 && guardX < len(lab)-1 && guardY < len(lab[0])-1 {
		prevX, prevY := guardX, guardY
		if string(lab[guardX][guardY]) == "^" && slices.Contains([]string{".", "X"}, string(lab[guardX-1][guardY])) {
			guardX--
		} else if string(lab[guardX][guardY]) == "v" && slices.Contains([]string{".", "X"}, string(lab[guardX+1][guardY])) {
			guardX++
		} else if string(lab[guardX][guardY]) == ">" && slices.Contains([]string{".", "X"}, string(lab[guardX][guardY+1])) {
			guardY++
		} else if string(lab[guardX][guardY]) == "<" && slices.Contains([]string{".", "X"}, string(lab[guardX][guardY-1])) {
			guardY--
		} else {
			view = rotateGuard(view)
			drawChar(lab, guardX, guardY, view)
			continue
		}
		drawChar(lab, guardX, guardY, view)
		drawChar(lab, prevX, prevY, 'X')
	}
}

func checkObstacleForward(lab []string, x int, y int, view rune, obstacle string) bool {
	switch view {
	case '^':
		return string(lab[x-1][y]) == obstacle
	case 'v':
		return string(lab[x+1][y]) == obstacle
	case '>':
		return string(lab[x][y+1]) == obstacle
	case '<':
		return string(lab[x][y-1]) == obstacle
	default:
		return false
	}
}

func checkCycle(lab []string) bool {
	guardX, guardY := findGuard(lab)
	var crashesMap map[string]int
	crashesMap = make(map[string]int)
	view := '^'
	for guardX > 0 && guardY > 0 && guardX < len(lab)-1 && guardY < len(lab[0])-1 {
		prevX, prevY := guardX, guardY
		if string(lab[guardX][guardY]) == "^" && slices.Contains([]string{".", "X"}, string(lab[guardX-1][guardY])) {
			guardX--
		} else if string(lab[guardX][guardY]) == "v" && slices.Contains([]string{".", "X"}, string(lab[guardX+1][guardY])) {
			guardX++
		} else if string(lab[guardX][guardY]) == ">" && slices.Contains([]string{".", "X"}, string(lab[guardX][guardY+1])) {
			guardY++
		} else if string(lab[guardX][guardY]) == "<" && slices.Contains([]string{".", "X"}, string(lab[guardX][guardY-1])) {
			guardY--
		} else {
			if checkObstacleForward(lab, guardX, guardY, view, "O") || checkObstacleForward(lab, guardX, guardY, view, "#") {
				key := string(guardX) + "," + string(guardY)
				crashes, ok := crashesMap[key]
				if !ok {
					crashesMap[key] = 1
				} else {
					crashesMap[key]++
					if crashes > 1 {
						return true
					}
				}
			}
			view = rotateGuard(view)
			drawChar(lab, guardX, guardY, view)
			continue
		}
		drawChar(lab, guardX, guardY, view)
		drawChar(lab, prevX, prevY, 'X')
	}

	return false
}

func findCycles(lab []string) int {
	cycles := 0
	tries := 0
	for i := 0; i < len(lab); i++ {
		for j := 0; j < len(lab[i]); j++ {
			if string(lab[i][j]) == "X" {
				newLab := append([]string(nil), lab...)
				drawChar(newLab, i, j, 'O')
				if checkCycle(newLab) {
					cycles++
				}
				tries++
			}
		}
	}
	return cycles
}

func main() {
	inputFile, _ := os.Open("day06/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	lab := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		lab = append(lab, scanner.Text())
	}

	// Part one
	positions := 0
	labPartOne := append([]string(nil), lab...)
	markPositions(labPartOne)
	positions = countX(labPartOne)
	fmt.Println("Distinct positions:", positions)

	// Find guard initial position and replace in the marked lab
	guardX, guardY := findGuard(lab)
	prevGuardX, prevGuardY := findGuard(labPartOne)
	drawChar(labPartOne, guardX, guardY, '^')
	drawChar(labPartOne, prevGuardX, prevGuardY, 'X')

	// Part two (reutilize X from part One)
	cycles := findCycles(labPartOne)
	fmt.Println("Cycles:", cycles)
}

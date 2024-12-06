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

func drawX(lab []string, x int, y int) []string {
	out := []rune(lab[x])
	out[y] = 'X'
	lab[x] = string(out)
	return lab
}

func drawLab(lab []string) {
	for i := 0; i < len(lab); i++ {
		fmt.Println(lab[i])
	}
}

func drawGuard(lab []string, x int, y int, view rune) []string {
	out := []rune(lab[x])
	out[y] = view
	lab[x] = string(out)
	return lab
}

func rotateGuard(view string) string {
	switch view {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
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

func markPositions(lab []string) []string {
	guardX, guardY := findGuard(lab)
	view := "^"
	for guardX > 0 && guardY > 0 && guardX < len(lab)-1 && guardY < len(lab[0])-1 {
		prevX, prevY := guardX, guardY
		if string(lab[guardX][guardY]) == "^" && slices.Contains([]string{".", "X"}, string(lab[guardX-1][guardY])) {
			view = "^"
			guardX--
		} else if string(lab[guardX][guardY]) == "v" && slices.Contains([]string{".", "X"}, string(lab[guardX+1][guardY])) {
			view = "v"
			guardX++
		} else if string(lab[guardX][guardY]) == ">" && slices.Contains([]string{".", "X"}, string(lab[guardX][guardY+1])) {
			view = ">"
			guardY++
		} else if string(lab[guardX][guardY]) == "<" && slices.Contains([]string{".", "X"}, string(lab[guardX][guardY-1])) {
			view = "<"
			guardY--
		} else {
			view = rotateGuard(view)
			lab = drawGuard(lab, guardX, guardY, rune(view[0]))
			continue
		}
		lab = drawGuard(lab, guardX, guardY, rune(view[0]))
		lab = drawX(lab, prevX, prevY)
	}

	return lab
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
	lab = markPositions(lab)

	positions = countX(lab)
	fmt.Println("Distinct positions:", positions)

	// Part two
}

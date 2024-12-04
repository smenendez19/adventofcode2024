package main

import (
	"bufio"
	"fmt"
	"os"
)

func findHorizontal(letterSoup []string, x int, y int, word string, reversed bool) bool {
	if !reversed {
		if y+len(word) > len(letterSoup[x]) {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			y++
		}
	} else {
		if y < len(word)-1 {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			y--
		}
	}

	return true
}

func findVertical(letterSoup []string, x int, y int, word string, reversed bool) bool {
	if !reversed {
		if x+len(word) > len(letterSoup) {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			x++
		}
	} else {
		if x < len(word)-1 {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			x--
		}
	}

	return true
}

func findDiagonalInf(letterSoup []string, x int, y int, word string, reversed bool) bool {
	if !reversed {
		if x+len(word) > len(letterSoup) || y+len(word) > len(letterSoup[x]) {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			x++
			y++
		}
	} else {
		if x+len(word) > len(letterSoup) || y < len(word)-1 {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			x++
			y--
		}
	}

	return true
}

func findDiagonalSup(letterSoup []string, x int, y int, word string, reversed bool) bool {
	if !reversed {
		if x < len(word)-1 || y+len(word) > len(letterSoup) {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			x--
			y++
		}
	} else {
		if x < len(word)-1 || y < len(word)-1 {
			return false
		}
		for _, c := range word {
			if string(letterSoup[x][y]) != string(c) {
				return false
			}
			x--
			y--
		}
	}

	return true
}

func findOcurrences(letterSoup []string, word string) int {
	ocurrences := 0

	for i := 0; i < len(letterSoup); i++ {
		for j := 0; j < len(letterSoup[i]); j++ {
			if findHorizontal(letterSoup, i, j, word, true) {
				ocurrences++
			}
			if findHorizontal(letterSoup, i, j, word, false) {
				ocurrences++
			}
			if findVertical(letterSoup, i, j, word, true) {
				ocurrences++
			}
			if findVertical(letterSoup, i, j, word, false) {
				ocurrences++
			}
			if findDiagonalInf(letterSoup, i, j, word, true) {
				ocurrences++
			}
			if findDiagonalInf(letterSoup, i, j, word, false) {
				ocurrences++
			}
			if findDiagonalSup(letterSoup, i, j, word, true) {
				ocurrences++
			}
			if findDiagonalSup(letterSoup, i, j, word, false) {
				ocurrences++
			}
		}
	}

	return ocurrences
}

func findCross(letterSoup []string) int {
	ocurrences := 0

	for i := 0; i < len(letterSoup); i++ {
		for j := 0; j < len(letterSoup[i]); j++ {
			// M M
			//  A
			// S S
			if findDiagonalSup(letterSoup, i, j, "AM", true) && findDiagonalInf(letterSoup, i, j, "AS", false) &&
				findDiagonalSup(letterSoup, i, j, "AM", false) && findDiagonalInf(letterSoup, i, j, "AS", true) {
				ocurrences++
			}
			// S M
			//  A
			// S M
			if findDiagonalSup(letterSoup, i, j, "AS", true) && findDiagonalInf(letterSoup, i, j, "AM", false) &&
				findDiagonalSup(letterSoup, i, j, "AM", false) && findDiagonalInf(letterSoup, i, j, "AS", true) {
				ocurrences++
			}
			// S S
			//  A
			// M M
			if findDiagonalSup(letterSoup, i, j, "AS", true) && findDiagonalInf(letterSoup, i, j, "AM", false) &&
				findDiagonalSup(letterSoup, i, j, "AS", false) && findDiagonalInf(letterSoup, i, j, "AM", true) {
				ocurrences++
			}
			// M S
			//  A
			// M S
			if findDiagonalSup(letterSoup, i, j, "AM", true) && findDiagonalInf(letterSoup, i, j, "AS", false) &&
				findDiagonalSup(letterSoup, i, j, "AS", false) && findDiagonalInf(letterSoup, i, j, "AM", true) {
				ocurrences++
			}

		}
	}

	return ocurrences
}

func main() {
	inputFile, _ := os.Open("day04/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	letterSoup := []string{}

	for scanner.Scan() {
		letterSoup = append(letterSoup, scanner.Text())
	}

	xmasOcurrences := findOcurrences(letterSoup, "XMAS")
	fmt.Println("XMAS:", xmasOcurrences)
	crossXMasOcurrences := findCross(letterSoup)
	fmt.Println("X-MAS:", crossXMasOcurrences)
}

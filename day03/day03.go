package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func partOne(line string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	results := re.FindAllString(line, -1)

	subTotal := 0

	for _, result := range results {
		matches := re.FindStringSubmatch(result)
		a, _ := strconv.Atoi(matches[1])
		b, _ := strconv.Atoi(matches[2])
		subTotal += (a * b)
	}

	return subTotal
}

func partTwo(line string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	results := re.FindAllString(line, -1)

	subTotal := 0
	enableMult := true

	for _, result := range results {
		if result == "don't()" {
			enableMult = false
			continue
		}
		if result == "do()" {
			enableMult = true
			continue
		}
		if enableMult {
			matches := re.FindStringSubmatch(result)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			subTotal += (a * b)
		}
	}

	return subTotal
}

func main() {
	inputFile, _ := os.Open("day03/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	corruptedMemory := ""

	totalMultPartOne := 0
	totalMultPartTwo := 0

	for scanner.Scan() {
		corruptedMemory += scanner.Text()
	}

	totalMultPartOne += partOne(corruptedMemory)
	totalMultPartTwo += partTwo(corruptedMemory)

	fmt.Println("Part One:", totalMultPartOne)
	fmt.Println("Part Two:", totalMultPartTwo)
}

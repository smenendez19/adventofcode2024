package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func everyEqual(array []string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] != array[0] {
			return false
		}
	}
	return true
}

func leftPad(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}

func concat(a, b int64) int64 {
	res := strconv.FormatInt(a, 10) + strconv.FormatInt(b, 10)
	result, _ := strconv.ParseInt(res, 10, 64)
	return result
}

func calculateSymbols(numbers []int64, symbols []string) int64 {
	var result int64
	result = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if symbols[i-1] == "0" {
			result = result + numbers[i]
		} else if symbols[i-1] == "1" {
			result = result * numbers[i]
		} else if symbols[i-1] == "2" {
			result = concat(result, numbers[i])
		}
	}
	return result
}

func checkSymbols(numbers []int64, result int64, symbols int) bool {
	binarySymbols := []string{}
	combinations := int(math.Round(math.Pow(float64(symbols), float64(len(numbers)))))
	for i := 0; i < combinations; i++ {
		symbols := leftPad(strconv.FormatInt(int64(i), symbols), len(numbers))
		binarySymbols = append(binarySymbols, symbols)
	}
	for i := 0; i < len(binarySymbols); i++ {
		if calculateSymbols(numbers, strings.Split(binarySymbols[i], "")) == result {
			return true
		}
	}
	return false
}

func calculatePartTwo(ecuations []string) int64 {
	var sum int64
	sum = 0
	for _, ecuation := range ecuations {
		result, _ := strconv.ParseInt(strings.Split(ecuation, ":")[0], 10, 64)
		numbersStr := strings.Split(strings.Split(ecuation, ":")[1], " ")[1:]
		numbers := []int64{}
		for _, number := range numbersStr {
			num, _ := strconv.ParseInt(number, 10, 64)
			numbers = append(numbers, num)
		}
		if checkSymbols(numbers, result, 3) {
			sum = sum + result
		}
	}
	return sum
}

func calculatePartOne(ecuations []string) int64 {
	var sum int64
	sum = 0
	for _, ecuation := range ecuations {
		result, _ := strconv.ParseInt(strings.Split(ecuation, ":")[0], 10, 64)
		numbersStr := strings.Split(strings.Split(ecuation, ":")[1], " ")[1:]
		numbers := []int64{}
		for _, number := range numbersStr {
			num, _ := strconv.ParseInt(number, 10, 64)
			numbers = append(numbers, num)
		}
		if checkSymbols(numbers, result, 2) {
			sum = sum + result
		}
	}
	return sum
}

func main() {
	inputFile, _ := os.Open("day07/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	ecuations := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		ecuations = append(ecuations, scanner.Text())
	}

	resultPartOne := calculatePartOne(ecuations)
	fmt.Println("Part one:", resultPartOne)

	resultPartTwo := calculatePartTwo(ecuations)
	fmt.Println("Part two:", resultPartTwo)
}

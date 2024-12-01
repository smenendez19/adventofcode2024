package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day01PartOne(leftList []int, rightList []int) int {
	totalDistance := 0

	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance *= -1
		}
		totalDistance += distance
	}

	return totalDistance
}

func day01PartTwo(leftList []int, rightList []int) int {
	similarityScore := 0

	idxRightList := 0
	for i := 0; i < len(leftList); i++ {
		for j := idxRightList; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				similarityScore += leftList[i]
			}
			if leftList[i] < rightList[j] {
				idxRightList = j
				break
			}
		}
	}

	return similarityScore
}

func main() {
	var leftList []int
	var rightList []int

	csvFile, _ := os.Open("day01/input.txt")
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		numbers := strings.Split(record[0], "   ")
		leftNum, _ := strconv.Atoi(numbers[0])
		rightNum, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	sort.Ints(leftList[:])
	sort.Ints(rightList[:])

	fmt.Println("Total distance:", day01PartOne(leftList, rightList))
	fmt.Println("Similarity score:", day01PartTwo(leftList, rightList))
}

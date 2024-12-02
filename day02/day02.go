package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Part one
func verifySafeReport(report []int) bool {
	diff := 0
	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] == 0 {
			return false
		} else if report[i]-report[i-1] > 3 || report[i]-report[i-1] < -3 {
			return false
		} else {
			if diff == 0 {
				diff = report[i] - report[i-1]
			} else if diff*(report[i]-report[i-1]) < 0 {
				return false
			}
		}
	}
	return true
}

// Part two

func verifySafeReportTolerance(report []int) bool {
	if verifySafeReport(report) {
		return true
	}

	// Check tolerance
	for i := 0; i < len(report); i++ {
		newReport := append([]int(nil), report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if verifySafeReport(newReport) {
			return true
		}
	}

	return false
}

func main() {
	csvFile, _ := os.Open("day02/input.txt")
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = ' '

	// Part one
	var safeReports = 0
	// Part two
	var safeTrulyReports = 0

	for {
		var report []int
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		for i := 0; i < len(record); i++ {
			num, _ := strconv.Atoi(record[i])
			report = append(report, num)
		}

		isSafe := verifySafeReport(report)
		if isSafe {
			safeReports++
		}

		isTrulySafe := verifySafeReportTolerance(report)
		if isTrulySafe {
			safeTrulyReports++
		}
	}

	fmt.Println("Safe reports:", safeReports)
	fmt.Println("Truly safe reports:", safeTrulyReports)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func find(list []string, element string) int {
	for i, item := range list {
		if item == element {
			return i
		}
	}
	return -1
}

func checkRules(page []string, pageOrderingRules []string) bool {
	for _, rule := range pageOrderingRules {
		rule := strings.Split(rule, "|")
		indexRuleBefore := find(page, rule[0])
		indexRuleAfter := find(page, rule[1])
		if indexRuleBefore == -1 || indexRuleAfter == -1 {
			continue
		}
		if indexRuleBefore > indexRuleAfter {
			return false
		}
	}
	return true
}

func reorderPage(page []string, pageOrderingRules []string) []string {

	for !checkRules(page, pageOrderingRules) {
		for _, rule := range pageOrderingRules {
			rule := strings.Split(rule, "|")
			indexRuleBefore := find(page, rule[0])
			indexRuleAfter := find(page, rule[1])
			if indexRuleBefore == -1 || indexRuleAfter == -1 {
				continue
			}
			if indexRuleBefore > indexRuleAfter {
				page = slices.Replace(page, indexRuleBefore, indexRuleBefore+1, rule[1])
				page = slices.Replace(page, indexRuleAfter, indexRuleAfter+1, rule[0])
			}
		}
	}

	return page
}

func main() {
	inputFile, _ := os.Open("day05/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	pageOrderingRules := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		pageOrderingRules = append(pageOrderingRules, scanner.Text())
	}

	pageNumbers := []string{}

	for scanner.Scan() {
		pageNumbers = append(pageNumbers, scanner.Text())
	}

	sumMiddleNumbers := 0

	incorrectOrderedPages := [][]string{}

	for _, page := range pageNumbers {
		pageList := strings.Split(page, ",")
		if checkRules(pageList, pageOrderingRules) {
			middleNumber, _ := strconv.Atoi(pageList[len(pageList)/2])
			sumMiddleNumbers += middleNumber
		} else {
			incorrectOrderedPages = append(incorrectOrderedPages, pageList)
		}
	}

	sumMiddleNumbersIncorrect := 0

	for _, page := range incorrectOrderedPages {
		orderedPage := reorderPage(page, pageOrderingRules)
		middleNumber, _ := strconv.Atoi(orderedPage[len(orderedPage)/2])
		sumMiddleNumbersIncorrect += middleNumber
	}

	fmt.Println("Sum:", sumMiddleNumbers)
	fmt.Println("Sum incorrect ordered pages:", sumMiddleNumbersIncorrect)
}

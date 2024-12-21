package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent_of_code_day_5() {
	file, ferr := os.Open("day_5_input.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	var beforeNums []int
	var afterNums []int
	var updates []string
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			fmt.Printf("ignore empty row\n")
			continue // continue between rules and updates
		}
		if splitRule := strings.Split(row, "|"); len(splitRule) == 2 { // format of updates
			beforeNum, err1 := strconv.Atoi(splitRule[0])
			if err1 != nil {
				panic(err1)
			}
			beforeNums = append(beforeNums, beforeNum)

			afterNum, err2 := strconv.Atoi(splitRule[1])
			if err2 != nil {
				panic(err2)
			}
			afterNums = append(afterNums, afterNum)
			fmt.Printf("add row %s to rules\n", row)
			continue
		}
		fmt.Printf("add row %s to updates\n", row)
		updates = append(updates, row) // everything else is an update
	}

	runningTotal := 0
	for i := 0; i < len(updates); i++ {
		updateStr := updates[i]
		update := parseUpdate(updateStr)
		if updatePassesAllRules(update, beforeNums, afterNums) {
			fmt.Printf("update %s passes all rules, so ignoring\n", updateStr)
		} else {
			reorderedUpdate := reorderUpdateToPassAllRules(update, beforeNums, afterNums)
			fmt.Printf("update %s does not pass all rules, but reordered ", updateStr)
			for _, value := range reorderedUpdate {
				fmt.Printf("%d,", value)
			}
			middlePageIndex := (len(reorderedUpdate) - 1) / 2
			middlePage := reorderedUpdate[middlePageIndex]
			fmt.Printf(" does, adds %d to running total %d\n", middlePage, runningTotal)
			runningTotal += middlePage
		}
	}
	fmt.Printf("Sum of middle page from correct updates %d\n", runningTotal)
}

func reorderUpdateToPassAllRules(input []int, beforeNums []int, afterNums []int) []int {
	if len(input) <= 1 {
		return input
	}
	currentSmallest := input[0]
	for i := 0; i < len(input); i++ {
		var comparisonArr []int
		comparisonArr = append(comparisonArr, currentSmallest, input[i])
		if !updatePassesAllRules(comparisonArr, beforeNums, afterNums) {
			currentSmallest = input[i]
		}
	}
	var inputWithoutSmallest []int
	for i := 0; i < len(input); i++ {
		if input[i] != currentSmallest {
			inputWithoutSmallest = append(inputWithoutSmallest, input[i])
		}
	}
	reordered := reorderUpdateToPassAllRules(inputWithoutSmallest, beforeNums, afterNums)

	var output []int
	output = append(output, currentSmallest)
	output = append(output, reordered...)
	return output
}

func parseUpdate(inputUpdate string) []int {
	splitUpdate := strings.Split(inputUpdate, ",")
	var result []int
	for i := 0; i < len(splitUpdate); i++ {
		update, err := strconv.Atoi(splitUpdate[i])
		if err != nil {
			panic(err)
		}
		result = append(result, update)
	}
	return result
}

func updatePassesAllRules(update []int, beforeNums []int, afterNums []int) bool {
	for j := 0; j < len(beforeNums); j++ {
		beforeNum := beforeNums[j]
		afterNum := afterNums[j]
		if updateFailsRule(update, beforeNum, afterNum) {
			return false
		}
	}
	return true
}

func updateFailsRule(update []int, beforeNum int, afterNum int) bool {
	beforeNumIndex := -1
	afterNumIndex := -1
	for i := 0; i < len(update); i++ {
		if update[i] == beforeNum {
			beforeNumIndex = i
		}
		if update[i] == afterNum {
			afterNumIndex = i
		}
	}
	if beforeNumIndex == -1 || afterNumIndex == -1 {
		return false
	}
	if afterNumIndex < beforeNumIndex {
		fmt.Printf("rule broke: %d should be before %d", beforeNum, afterNum)
		return true
	}
	return false
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func advent_of_code_day_2() {
	file, ferr := os.Open("day_2_input.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	safeLevelsCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		levelsStrArr := strings.Split(line, " ")
		var levelsIntArr []int
		for _, str := range levelsStrArr {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			levelsIntArr = append(levelsIntArr, num)
		}
		isLevelSafe := isLevelSafeWithDampener(levelsIntArr)
		if isLevelSafe {
			safeLevelsCount++
			fmt.Printf("%s level safe!\n", line)
		} else {
			fmt.Printf("%s level not safe!\n", line)
		}
	}
	fmt.Printf("Safe levels count: %d\n", safeLevelsCount)
}

func isLevelSafe(reports []int) bool {
	length := len(reports)
	if length <= 1 {
		return true
	}
	initialDirection := math.Signbit(float64(reports[0] - reports[1]))
	for i := 0; i < length-1; i++ {
		change := float64(reports[i] - reports[i+1])
		if (math.Abs(change) < 1) || (math.Abs(change) > 3) {
			return false // change is too large!
		}
		direction := math.Signbit(change)
		if direction != initialDirection {
			return false // the direction changed!
		}
	}

	return true
}

func isLevelSafeWithDampener(reports []int) bool {
	if isLevelSafe(reports) {
		return true
	}
	for i := 0; i < len(reports); i++ {
		if isLevelSafe(remove(reports, i)) {
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	newArr := make([]int, len(slice)-1)
	for i := 0; i < len(slice)-1; i++ {
		if i < s {
			newArr[i] = slice[i]
		} else {
			newArr[i] = slice[i+1]
		}
	}
	return newArr
}

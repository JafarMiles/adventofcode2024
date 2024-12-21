package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

func advent_of_code_day_4() {
	file, ferr := os.Open("day_4_input.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	var rows []string
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
		println(row)
	}

	runningTotal := 0
	for i := 0; i < len(rows)-2; i++ {
		for j := 0; j < len(rows[0])-2; j++ {

			diagonal := string(getChar(rows[i], j)) + string(getChar(rows[i+1], j+1)) + string(getChar(rows[i+2], j+2))
			otherDiagonal := string(getChar(rows[i+2], j)) + string(getChar(rows[i+1], j+1)) + string(getChar(rows[i], j+2))

			if matchesWord(diagonal) && matchesWord(otherDiagonal) {
				runningTotal++
			}
		}
	}
	fmt.Printf("Found %d matches\n", runningTotal)
}

func matchesWord(candidate string) bool {
	return candidate == "MAS" || candidate == "SAM"
}

func advent_of_code_day_4_part_1(rows []string) {
	runningTotal := 0
	foundInRows := countMatches(rows)
	fmt.Printf("Found %d matches in rows\n", foundInRows)
	runningTotal += foundInRows

	columns := rotate(rows)

	foundInColumns := countMatches(columns)
	fmt.Printf("Found %d matches in columns\n", foundInColumns)
	runningTotal += foundInColumns

	forwardDiagonals := converToDiagonal(rows)
	foundInDiagonals := countMatches(forwardDiagonals)
	fmt.Printf("Found %d matches in diagonal, bottom left to top right\n", foundInDiagonals)
	runningTotal += foundInDiagonals

	backwardsDiagonals := converToDiagonal(columns)
	runningTotal += countMatches(backwardsDiagonals)
	fmt.Printf("Found %d matches in other diagonal, bottom left to top right\n", foundInDiagonals)

	fmt.Printf("found %d XMAS in the puzzle\n", runningTotal)
}

func converToDiagonal(rows []string) []string {
	var diagonals []string
	for i := 0; i < len(rows[0])+len(rows)-1; i++ {
		diagonal := ""
		startRow := int(math.Min(float64(i), float64(len(rows)-1)))
		startColumn := int(math.Max(0, float64(i-len(rows)+1)))

		for j := 0; j < startRow-startColumn+1; j++ {
			row := startRow - j
			column := startColumn + j

			diagonal += string(getChar(rows[row], column))
		}
		fmt.Printf("row %d col %d: %s\n", startRow, startColumn, diagonal)
		diagonals = append(diagonals, diagonal)
	}
	return diagonals
}

func countMatches(lines []string) int {
	xmasRegex, rErr := regexp.Compile("XMAS")
	if rErr != nil {
		panic(rErr)
	}

	reverseXmasRegex, rrErr := regexp.Compile("SAMX")
	if rrErr != nil {
		panic(rrErr)
	}

	runningTotal := 0
	for i := 0; i < len(lines); i++ {
		matches := xmasRegex.FindAllString(lines[i], -1)
		runningTotal += len(matches)

		reverseMatches := reverseXmasRegex.FindAllString(lines[i], -1)
		runningTotal += len(reverseMatches)
	}
	return runningTotal
}

func getChar(str string, index int) rune {
	return []rune(str)[index]
}

func rotate(rows []string) []string {
	var columns []string
	for i := 0; i < len(rows[0]); i++ {
		column := ""
		for j := 0; j < len(rows); j++ {
			column += string(getChar(rows[j], len(rows[0])-i-1))
		}
		columns = append(columns, column)
		println(column)
	}
	return columns
}

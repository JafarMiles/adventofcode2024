package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	Row, Column int
}

func advent_of_code_day_8() { // part 1
	file, ferr := os.Open("day_8.in")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	// create a set - we will put the anti-nodes here
	mapOfFrequenciesToCoords := make(map[rune][]Coordinate)

	row := 0
	colLength := 0
	for scanner.Scan() {
		line := scanner.Text()
		colLength = len(line)
		for col := 0; col < len(line); col++ {
			character := rune(line[col])
			if character != '.' {
				coordinates := mapOfFrequenciesToCoords[character]
				coord := Coordinate{row, col}
				coordinates = append(coordinates, coord)
				mapOfFrequenciesToCoords[character] = coordinates
			}
		}
		row++
	}
	rowLength := row
	setOfAntiNodes := make(map[Coordinate]void)
	count := 0
	for frequencies := range mapOfFrequenciesToCoords {
		coordinates := mapOfFrequenciesToCoords[frequencies]
		for i := 0; i < len(coordinates); i++ {
			for j := i + 1; j < len(coordinates); j++ {
				first := coordinates[i]
				second := coordinates[j]

				antiNotes := find_anti_nodes_part_2(first, second, rowLength, colLength)

				for _, antiNode := range antiNotes {
					if coordinate_in_bounds(antiNode, rowLength, colLength) {
						_, ok := setOfAntiNodes[antiNode]
						if !ok {
							count++
							setOfAntiNodes[antiNode] = member
						}
					}
				}
			}
		}
	}
	fmt.Printf("number of anti-nodes is: %d\n", count)
}

func find_anti_nodes_part_2(first Coordinate, second Coordinate, rowLength int, colLength int) []Coordinate {
	var antiNodes []Coordinate
	diffRow := second.Row - first.Row
	diffCol := second.Column - first.Column

	// find the highest common factor between diffRow and diffCol
	hcf := find_highest_common_factor(diffRow, diffCol)
	smallestRowDiff := diffRow / hcf
	smallestColDiff := diffCol / hcf

	for i := 0; true; i++ {
		antiNoteCandidate := Coordinate{
			first.Row - smallestRowDiff*i,
			first.Column - smallestColDiff*i,
		}
		if coordinate_in_bounds(antiNoteCandidate, rowLength, colLength) {
			antiNodes = append(antiNodes, antiNoteCandidate)
		} else {
			break
		}
	}

	for i := 0; true; i++ {
		antiNoteCandidate := Coordinate{
			first.Row + smallestRowDiff*i,
			first.Column + smallestColDiff*i,
		}
		if coordinate_in_bounds(antiNoteCandidate, rowLength, colLength) {
			antiNodes = append(antiNodes, antiNoteCandidate)
		} else {
			break
		}
	}
	return antiNodes
}

func find_highest_common_factor(first int, second int) int {
	if first < 0 {
		return find_highest_common_factor(first*-1, second)
	}
	if second < 0 {
		return find_highest_common_factor(first, second*-1)
	}
	if second > first {
		return find_highest_common_factor(second, first)
	}
	if first%second == 0 {
		return second
	}
	return find_highest_common_factor(first-second, second)
}

func find_anti_nodes_part_1(first Coordinate, second Coordinate, rowLength int, colLength int) []Coordinate {
	var antiNodes []Coordinate

	diffRow := second.Row - first.Row
	diffCol := second.Column - first.Column

	secondSideAntiNode := Coordinate{
		second.Row + diffRow,
		second.Column + diffCol,
	}
	if coordinate_in_bounds(secondSideAntiNode, rowLength, colLength) {
		antiNodes = append(antiNodes, secondSideAntiNode)
	}

	firstSideAntiNode := Coordinate{
		first.Row - diffRow,
		first.Column - diffCol,
	}
	if coordinate_in_bounds(firstSideAntiNode, rowLength, colLength) {
		antiNodes = append(antiNodes, firstSideAntiNode)
	}
	return antiNodes
}

func coordinate_in_bounds(coord Coordinate, rowLength int, colLength int) bool {
	return coord.Row >= 0 && coord.Row < rowLength && coord.Column >= 0 && coord.Column < colLength
}

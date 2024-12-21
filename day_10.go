package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent_of_code_day_10() { // part 1
	file, ferr := os.Open("day_10.in")
	if ferr != nil {
		panic(ferr)
	}

	var trailHeads []Coordinate
	var topographicalMap [][]int
	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for columnIndex, character := range line {
			height := int(rune(character)) - int('0')
			row = append(row, height)
			if height == 0 {
				trailHead := Coordinate{rowIndex, columnIndex}
				trailHeads = append(trailHeads, trailHead)
			}
		}
		topographicalMap = append(topographicalMap, row)
		rowIndex++
	}

	totalScore := 0
	for _, trailHead := range trailHeads {
		totalScore += rating(topographicalMap, trailHead)
	}
	fmt.Printf("total score of trail heads %d\n", totalScore)
}

func rating(topographicalMap [][]int, trailHead Coordinate) int {
	currentLevel := make(map[Coordinate]int)
	currentLevel[trailHead] = 1
	for i := 1; i <= 9; i++ {
		currentLevel = find_next_steps_in_hiking_trail(topographicalMap, i, currentLevel)
	}
	totalRating := 0
	for _, countOfPathsToCoordinate := range currentLevel {
		totalRating += countOfPathsToCoordinate
	}
	return totalRating
}

func score(topographicalMap [][]int, trailHead Coordinate) int {
	currentLevel := make(map[Coordinate]int)
	currentLevel[trailHead] = 1
	for i := 1; i <= 9; i++ {
		currentLevel = find_next_steps_in_hiking_trail(topographicalMap, i, currentLevel)
	}
	return len(currentLevel)
}

func find_next_steps_in_hiking_trail(topographicalMap [][]int, level int, startingPoints map[Coordinate]int) map[Coordinate]int {
	nextLevel := make(map[Coordinate]int)
	rowLength := len(topographicalMap)
	colLength := len(topographicalMap[0])
	for startingPoint, pathsToStartingPoint := range startingPoints {
		neighbours := find_neighbours(startingPoint, rowLength, colLength)
		for _, neighbour := range neighbours {
			if topographicalMap[neighbour.Row][neighbour.Column] == level {
				count := nextLevel[neighbour]
				nextLevel[neighbour] = count + pathsToStartingPoint
			}
		}
	}
	return nextLevel
}

func find_neighbours(startingPoint Coordinate, rowLength int, colLength int) []Coordinate {
	var neighbours []Coordinate
	up := Coordinate{startingPoint.Row - 1, startingPoint.Column}
	if is_in_bound(up, rowLength, colLength) {
		neighbours = append(neighbours, up)
	}
	right := Coordinate{startingPoint.Row, startingPoint.Column + 1}
	if is_in_bound(right, rowLength, colLength) {
		neighbours = append(neighbours, right)
	}
	//down
	down := Coordinate{startingPoint.Row + 1, startingPoint.Column}
	if is_in_bound(down, rowLength, colLength) {
		neighbours = append(neighbours, down)
	}
	//left
	left := Coordinate{startingPoint.Row, startingPoint.Column - 1}
	if is_in_bound(left, rowLength, colLength) {
		neighbours = append(neighbours, left)
	}
	return neighbours
}

func is_in_bound(point Coordinate, rowLength int, colLength int) bool {
	return point.Row >= 0 && point.Row < rowLength && point.Column >= 0 && point.Column < colLength
}

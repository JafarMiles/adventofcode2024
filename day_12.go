package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func advent_of_code_day_12() { // part 1
	file, ferr := os.Open("day_12.in")
	if ferr != nil {
		panic(ferr)
	}

	var mapOfGarden []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		mapOfGarden = append(mapOfGarden, text)
	}
	// find region, and perimeter
	//  choose point, find neighbours, calculate boundary found so far, repeat
	// repeat
	// sum total
	locationsVisited := make(map[Coordinate]void)
	total_cost := 0

	for rowIndex := 0; rowIndex < len(mapOfGarden); rowIndex++ {
		for colIndex := 0; colIndex < len(mapOfGarden[rowIndex]); colIndex++ {
			location := Coordinate{rowIndex, colIndex}
			_, ok := locationsVisited[location]
			if ok {
				continue // already visited
			}
			cost := calculate_fence_cost(mapOfGarden, location, locationsVisited)
			total_cost += cost
		}
	}
	fmt.Printf("the total cost of fencing is %d\n", total_cost)
}

func calculate_fence_cost(mapOfGarden []string, location Coordinate, locationsVisited map[Coordinate]void) int {
	locationsInRegion := make(map[Coordinate]void)
	leftBoundary := make(map[Coordinate]void)
	upBoundary := make(map[Coordinate]void)
	rightBoundary := make(map[Coordinate]void)
	downBoundary := make(map[Coordinate]void)

	area, _ := add_location_to_region(mapOfGarden, location, locationsVisited, locationsInRegion, 0, 0, leftBoundary, upBoundary, rightBoundary, downBoundary)

	number_of_sides := calculate_number_of_sides(leftBoundary, upBoundary, rightBoundary, downBoundary)

	region := rune(mapOfGarden[location.Row][location.Column])
	fmt.Printf("calculate_fence_cost: region %s, area %d * number of side %d = %d\n", string(region), area, number_of_sides, area*number_of_sides)

	return area * number_of_sides
}

func calculate_number_of_sides(leftBoundary map[Coordinate]void, upBoundary map[Coordinate]void, rightBoundary map[Coordinate]void, downBoundary map[Coordinate]void) int {
	left_sides := calculate_number_of_vertical_sides(leftBoundary)
	right_sides := calculate_number_of_vertical_sides(rightBoundary)
	up_sides := calculate_number_of_horizontal_sides(upBoundary)
	down_sides := calculate_number_of_horizontal_sides(downBoundary)
	fmt.Printf("%d left sides, %d right sides, %d up sides, %d down side\n", left_sides, right_sides, up_sides, down_sides)
	return left_sides + right_sides + up_sides + down_sides
}

func calculate_number_of_horizontal_sides(horizontalBoundary map[Coordinate]void) int {
	row_edges := make(map[int][]int)
	for location := range horizontalBoundary {
		row_edges[location.Row] = append(row_edges[location.Row], location.Column)
	}
	totalSides := 0

	for _, segments := range row_edges {
		slices.Sort(segments)
		sides := 1
		previous := segments[0]
		for i := 1; i < len(segments); i++ {
			if segments[i]-previous != 1 {
				sides++ // not the same side
			}
			previous = segments[i]
		}
		totalSides += sides
	}
	return totalSides
}

func calculate_number_of_vertical_sides(verticalBoundary map[Coordinate]void) int {
	column_edges := make(map[int][]int)
	for location := range verticalBoundary {
		column_edges[location.Column] = append(column_edges[location.Column], location.Row)
	}
	totalSides := 0

	for _, segments := range column_edges {
		slices.Sort(segments)
		sides := 1
		previous := segments[0]
		for i := 1; i < len(segments); i++ {
			if segments[i]-previous != 1 {
				sides++ // not the same side
			}
			previous = segments[i]
		}
		totalSides += sides
	}
	return totalSides
}

func add_location_to_region(mapOfGarden []string, location Coordinate, locationsVisited map[Coordinate]void, locationsInRegion map[Coordinate]void, area int, perimeter int, leftBoundary map[Coordinate]void, upBoundary map[Coordinate]void, rightBoundary map[Coordinate]void, downBoundary map[Coordinate]void) (int, int) {
	locationsVisited[location] = member
	_, ok := locationsInRegion[location]
	if !ok {
		locationsInRegion[location] = member
		area++
		neighbours := findNeighbours(mapOfGarden, location, leftBoundary, upBoundary, rightBoundary, downBoundary)
		perimeter += 4 - len(neighbours)

		for _, neighbour := range neighbours {
			newArea, newPerimeter := add_location_to_region(mapOfGarden, neighbour, locationsVisited, locationsInRegion, area, perimeter, leftBoundary, upBoundary, rightBoundary, downBoundary)
			area = newArea
			perimeter = newPerimeter
		}
	}
	return area, perimeter
}

func findNeighbours(mapOfGarden []string, location Coordinate, leftBoundary map[Coordinate]void, upBoundary map[Coordinate]void, rightBoundary map[Coordinate]void, downBoundary map[Coordinate]void) []Coordinate {
	up := Coordinate{location.Row - 1, location.Column}
	down := Coordinate{location.Row + 1, location.Column}
	left := Coordinate{location.Row, location.Column - 1}
	right := Coordinate{location.Row, location.Column + 1}

	region := rune(mapOfGarden[location.Row][location.Column])

	var neighbours []Coordinate
	if location_is_in_region(mapOfGarden, up, region) {
		neighbours = append(neighbours, up)
	} else {
		upBoundary[location] = member
	}
	if location_is_in_region(mapOfGarden, down, region) {
		neighbours = append(neighbours, down)
	} else {
		downBoundary[location] = member
	}
	if location_is_in_region(mapOfGarden, left, region) {
		neighbours = append(neighbours, left)
	} else {
		leftBoundary[location] = member
	}
	if location_is_in_region(mapOfGarden, right, region) {
		neighbours = append(neighbours, right)
	} else {
		rightBoundary[location] = member
	}
	return neighbours
}

func location_is_in_region(mapOfGarden []string, location Coordinate, region rune) bool {
	rowLength := len(mapOfGarden)
	colLength := len(mapOfGarden[0])
	if location.Row < 0 || location.Row >= rowLength {
		return false
	}
	if location.Column < 0 || location.Column >= colLength {
		return false
	}
	return rune(mapOfGarden[location.Row][location.Column]) == region
}

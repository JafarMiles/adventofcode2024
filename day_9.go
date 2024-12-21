package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent_of_code_day_9() { // part 1
	file, ferr := os.Open("day_9.in")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	emptySpace := -1

	var diskArray []int
	fileId := 0
	for index, character := range input {
		numberOfSpaces := int(rune(character)) - int('0')
		valueToSet := emptySpace

		isFileBlock := index%2 == 0
		if isFileBlock {
			valueToSet = fileId
			fileId++
		}

		for i := 0; i < numberOfSpaces; i++ {
			diskArray = append(diskArray, valueToSet)
		}
	}

	diskArray = compart_part2(diskArray, emptySpace)

	checksum := compute_checksum(diskArray, emptySpace)
	fmt.Printf("checksum for the compacted disk array is %d\n", checksum)
}

func compart_part2(diskArray []int, emptySpace int) []int {
	fileLength := 1
	previousFileId := diskArray[len(diskArray)-1]
	previousFileIndex := len(diskArray) - 1
	for i := len(diskArray) - 2; i >= 0; i-- {
		nextFileId := diskArray[i]
		if nextFileId == emptySpace {
			continue // ignore empty spaces
		}

		if previousFileId == nextFileId {
			fileLength++
			previousFileIndex = i
			continue //same as last one so don't move anything yet
		}

		// try to move the block if we can!
		diskArray = try_move(diskArray, previousFileId, previousFileIndex, fileLength, emptySpace)

		previousFileId = nextFileId
		fileLength = 1
		previousFileIndex = i
	}
	return diskArray
}

func try_move(diskArray []int, fileId int, indexOfBeginningOfFile int, fileLength int, emptySpace int) []int {
	freeSpace := 0
	for j := 0; j < indexOfBeginningOfFile; j++ {
		if diskArray[j] == emptySpace {
			freeSpace++
		} else {
			freeSpace = 0
		}
		if freeSpace == fileLength {
			return move(diskArray, j-fileLength+1, indexOfBeginningOfFile, fileLength, fileId, emptySpace)
		}
	}
	return diskArray
}

func move(diskArray []int, indexBeginningOfFreeSpace int, indexOfBeginningOfFile int, fileLength int, fileId int, emptySpace int) []int {
	//fmt.Printf("Attempt to compact %d to index %d, with freespace %d\n", fileId, indexBeginningOfFreeSpace, fileLength)
	//fmt.Printf("DiskArray (before move): ")
	//print_disk_array(diskArray, emptySpace)
	//fmt.Printf("\n")
	for k := 0; k < fileLength; k++ {
		diskArray[indexBeginningOfFreeSpace+k] = fileId
		diskArray[indexOfBeginningOfFile+k] = emptySpace
	}
	//fmt.Printf("DiskArray (after move):  ")
	//print_disk_array(diskArray, emptySpace)
	//fmt.Printf("\n")
	return diskArray
}

func print_disk_array(diskArray []int, emptySpace int) {
	for x := 0; x < len(diskArray); x++ {
		if diskArray[x] == emptySpace {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", diskArray[x])
		}
	}
}

func compart_part1(diskArray []int, emptySpace int) []int {
	for j := 0; j < len(diskArray); j++ {
		diskArray = compact(diskArray, emptySpace, j)
	}
	return diskArray
}

func compact(diskArray []int, empty int, index int) []int {
	if index >= len(diskArray) {
		return diskArray
	}
	if diskArray[index] != empty {
		return diskArray
	}
	lastItem := diskArray[len(diskArray)-1]
	if lastItem == empty {
		return compact(diskArray[0:len(diskArray)-1], empty, index)
	}
	diskArray[index] = lastItem
	return diskArray[0 : len(diskArray)-1]
}

func compute_checksum(diskArray []int, empty int) int {
	checkSum := 0
	for index, fileId := range diskArray {
		if fileId != empty {
			checkSum += index * fileId
		}
	}
	return checkSum
}

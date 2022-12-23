package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const NORTH = "north"
const SOUTH = "south"
const WEST = "west"
const EAST = "east"

const NOTHING = "nothing"

type position [2]int

var elfs = map[position]bool{}

var directionsMap = map[string][]position{
	"north": {{-1, 0}, {-1, 1}, {-1, -1}},
	"south": {{1, 0}, {1, 1}, {1, -1}},
	"west":  {{0, -1}, {-1, -1}, {1, -1}},
	"east":  {{0, 1}, {-1, 1}, {1, 1}},
}

var directions = []string{NORTH, SOUTH, WEST, EAST}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	for index, line := range lines {
		for i, _ := range line {
			item := string(line[i])

			if item == "#" {
				elfs[position{index, i}] = true
			}
		}
	}

	// solvePart1()
	solvePart2()
}

func solvePart1() {
	rounds := 10

	for round := 0; round < rounds; round++ {
		proposedPositions := getProposedPositions(round)
		if len(proposedPositions) == 0 { // no more elfs to move
			return
		}

		move(proposedPositions)
	}

	calculateResult()
}

func solvePart2() {
	round := 0

	for {
		proposedPositions := getProposedPositions(round)
		if len(proposedPositions) == 0 {
			break
		}

		move(proposedPositions)

		round++
	}

	fmt.Println(round + 1)
}

func calculateResult() {
	minRow := math.MaxInt
	maxRow := 0

	minCol := math.MaxInt
	maxCol := 0

	for elfPosition, _ := range elfs {
		row := elfPosition[0]
		col := elfPosition[1]

		if row < minRow {
			minRow = row
		}

		if row > maxRow {
			maxRow = row
		}

		if col < minCol {
			minCol = col
		}

		if col > maxCol {
			maxCol = col
		}
	}

	count := 0

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if !isElfPosition(position{row, col}) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func move(proposedPositions map[position][]position) {
	for newPosition, elfPositios := range proposedPositions {
		if len(elfPositios) == 1 { // the elfs cannot move if len is greater than 1
			elfPosition := elfPositios[0]

			delete(elfs, elfPosition)
			elfs[newPosition] = true
		}
	}
}

func getProposedPositions(directionsMapStartIndex int) map[position][]position {
	proposedPositions := map[position]string{}

	for elfPosition := range elfs {
		if !canElfDoSomething(elfPosition) {
			proposedPositions[elfPosition] = NOTHING
			continue
		}

		directionsIndex := directionsMapStartIndex

		for i := 0; i < 4; i++ {
			directionName := getDirectionName(directionsIndex)
			if canElfMoveToDirection(directionName, elfPosition) {
				proposedPositions[elfPosition] = directionName
				break
			}

			directionsIndex++
		}
	}

	result := map[position][]position{}

	for elfPosition, directionName := range proposedPositions {
		if directionName == NOTHING {
			continue
		}

		proposedPosition := getNewPosition(directionsMap[directionName][0], elfPosition)

		if _, ok := result[proposedPosition]; !ok {
			result[proposedPosition] = []position{elfPosition}
		} else {
			elfsArray := result[proposedPosition]
			elfsArray = append(elfsArray, elfPosition)
			result[proposedPosition] = elfsArray
		}
	}

	return result
}

func getDirectionName(index int) string {
	return directions[index%len(directions)]
}

func getNewPosition(direction position, pos position) position {
	return position{direction[0] + pos[0], direction[1] + pos[1]}
}

func canElfDoSomething(elfPosition position) bool {
	directions := []position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for _, direction := range directions {
		if isElfPosition(getNewPosition(direction, elfPosition)) {
			return true
		}
	}

	return false
}

func canElfMoveToDirection(directionName string, elfPosition position) bool {
	directions := directionsMap[directionName]

	for _, dirPosition := range directions {
		newPosition := getNewPosition(dirPosition, elfPosition)
		if isElfPosition(newPosition) {
			return false
		}
	}

	return true
}

func isElfPosition(pos position) bool {
	return elfs[pos]
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position [2]int

var walls = map[position]bool{}
var opened = map[position]bool{}

var right = position{0, 1}
var left = position{0, -1}
var up = position{-1, 0}
var down = position{1, 0}

// (x, y) -> (y, -x)
var clockwiseTurn = map[position]position{
	right: down,
	left:  up,
	up:    right,
	down:  left,
}

// (x, y) -> (-y, x)
var counterclockwiseTurn = map[position]position{
	right: up,
	left:  down,
	up:    left,
	down:  right,
}

var directionsWeight = map[position]int{
	right: 0,
	down:  1,
	left:  2,
	up:    3,
}

var currentDirection = right
var currentPosition = position{}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	path := ""

	isCurrentPositionInitialized := false

	for index, line := range lines {
		for i := range line {
			item := string(line[i])

			if item == "#" {
				walls[position{index, i}] = true
			} else if item == "." {
				opened[position{index, i}] = true

				if index == 0 && !isCurrentPositionInitialized {
					currentPosition = position{index, i}
					isCurrentPositionInitialized = true
				}
			}
		}

		if line == "" {
			path = lines[index+1]
			break
		}
	}

	moves := processPath(path)
	for _, next := range moves {
		if next == "R" || next == "L" {
			turn(next)
		} else {
			move(next)
		}
	}

	fmt.Println(calculateResult())
}

func calculateResult() int {
	return 1000*(currentPosition[0]+1) + 4*(currentPosition[1]+1) + directionsWeight[currentDirection]
}

func turn(direction string) {
	if direction == "R" {
		currentDirection = clockwiseTurn[currentDirection]
	}

	if direction == "L" {
		currentDirection = counterclockwiseTurn[currentDirection]
	}
}

func move(str string) {
	num, _ := strconv.Atoi(str)

	for i := 0; i < num; i++ {
		newPosition := getNewPosition()
		if isWall(newPosition) { // cannot move anymore
			break
		}

		if isOpen(newPosition) { // move one step
			currentPosition = newPosition
			continue
		}

		// we're out of the board here and need to wrap around to the other side of the board
		prevPosition := newPosition
		reverseDirection()

		for {
			newPosition := position{currentDirection[0] + prevPosition[0], currentDirection[1] + prevPosition[1]}
			if isOutOfBoard(newPosition) {
				if isWall(prevPosition) {
					reverseDirection()
					break
				}

				currentPosition = prevPosition
				reverseDirection()
				break
			} else {
				prevPosition = newPosition
			}
		}
	}
}

func isWall(pos position) bool {
	return walls[pos]
}

func isOpen(pos position) bool {
	return opened[pos]
}

func isOutOfBoard(pos position) bool {
	return !isWall(pos) && !isOpen(pos)
}

func getNewPosition() position {
	return position{currentPosition[0] + currentDirection[0], currentPosition[1] + currentDirection[1]}
}

func reverseDirection() {
	if currentDirection == left {
		currentDirection = right
	} else if currentDirection == right {
		currentDirection = left
	} else if currentDirection == up {
		currentDirection = down
	} else if currentDirection == down {
		currentDirection = up
	}
}

func processPath(path string) []string {
	result := []string{}

	current := ""
	for i := 0; i < len(path); i++ {
		if path[i] == 'L' || path[i] == 'R' {
			result = append(result, current)
			result = append(result, string(path[i]))
			current = ""
		} else {
			current += string(path[i])
		}
	}

	result = append(result, current)

	return result
}

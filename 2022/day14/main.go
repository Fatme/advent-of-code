package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const SAND = "o"
const ROCK = "#"

const rows = 200
const cols = 560

var board = [rows][cols]string{}

var startX = 500
var startY = 0

func drawLine(x1, y1, x2, y2 int) {
	if x1 == x2 {
		min := int(math.Min(float64(y1), float64(y2)))
		max := int(math.Max(float64(y1), float64(y2)))

		for i := min; i <= max; i++ {
			markRock(x1, i)
		}
	}

	if y1 == y2 {
		min := int(math.Min(float64(x1), float64(x2)))
		max := int(math.Max(float64(x1), float64(x2)))

		for i := min; i <= max; i++ {
			markRock(i, y1)
		}
	}
}

func move(x, y int) bool {
	for y < rows-1 {
		if isPositionFree(x, y+1) {
			y++
		} else if isPositionFree(x-1, y+1) {
			x--
			y++
		} else if isPositionFree(x+1, y+1) {
			x++
			y++
		} else {
			markSand(x, y)
			break
		}
	}

	return y != rows-1
}

func isPositionFree(x, y int) bool {
	return isPositionValid(x, y) && board[y][x] != ROCK && board[y][x] != SAND
}

func isPositionValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < cols && y < rows
}

func markSand(x, y int) {
	if isPositionValid(x, y) {
		board[y][x] = SAND
	}
}

func markRock(x, y int) {
	if isPositionValid(x, y) {
		board[y][x] = ROCK
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	board[startY][startX] = "+"

	for _, line := range lines {
		pairs := strings.Split(line, " -> ")

		previousX := 0
		previousY := 0

		for pairIndex, pair := range pairs {
			parts := strings.Split(pair, ",")

			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			if pairIndex == 0 {
				board[y][x] = ROCK
			} else {
				drawLine(x, y, previousX, previousY)
			}

			previousX = x
			previousY = y
		}
	}

	index := 0
	for index >= 0 {
		if !move(startX, startY) {
			fmt.Println(index)
			break
		}
		index++
	}
}

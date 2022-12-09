package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func (p *Position) id() string {
	return fmt.Sprintf("%d->%d", p.x, p.y)
}

func (p *Position) distance(other *Position) int {
	return int(math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y)))
}

func (p *Position) move(xSteps int, ySteps int) *Position {
	return &Position{x: p.x + xSteps, y: p.y + ySteps}
}

type Rope struct {
	head                 *Position
	tail                 *Position
	visitedTailPositions map[string]*Position
}

func (r *Rope) moveHead(direction string) {
	if direction == "R" {
		r.head = r.head.move(0, 1)
	}

	if direction == "L" {
		r.head = r.head.move(0, -1)
	}

	if direction == "U" {
		r.head = r.head.move(-1, 0)
	}

	if direction == "D" {
		r.head = r.head.move(1, 0)
	}
}

func (r *Rope) moveTail() {
	min := math.MaxInt
	result := &Position{}

	tail := r.tail
	head := r.head

	// (-1, -1), (-1, 0), (-1, 1),
	// (0, -1), (0, 0), (0, 1),
	// (1, -1), (1, 0), (1, 1)
	possiblePositions := []*Position{
		{x: tail.x - 1, y: tail.y - 1},
		{x: tail.x - 1, y: tail.y},
		{x: tail.x - 1, y: tail.y + 1},

		{x: tail.x, y: tail.y - 1},
		{x: tail.x, y: tail.y},
		{x: tail.x, y: tail.y + 1},

		{x: tail.x + 1, y: tail.y - 1},
		{x: tail.x + 1, y: tail.y},
		{x: tail.x + 1, y: tail.y + 1},
	}

	for _, position := range possiblePositions {
		distance := head.distance(position)
		if distance < min {
			min = distance
			result = position
		}
	}

	r.tail = result

	r.visitedTailPositions[r.tail.id()] = tail
}

func (r *Rope) shouldMoveTail() bool {
	return math.Abs(float64(r.head.x-r.tail.x)) > 1 || math.Abs(float64(r.head.y-r.tail.y)) > 1
}

func main() {
	var rope = &Rope{head: &Position{0, 0}, tail: &Position{0, 0}, visitedTailPositions: map[string]*Position{}}

	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		for i := 0; i < steps; i++ {
			rope.moveHead(direction)
			if rope.shouldMoveTail() {
				rope.moveTail()
			}
		}
	}

	fmt.Println(len(rope.visitedTailPositions))
}

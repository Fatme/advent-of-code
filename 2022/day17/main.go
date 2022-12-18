package main

import "fmt"

const ROCK = '#'
const EMPTY = '.'

const rocksCount = 2022
const maxRockHeight = 4

const towerCols = 7
const towerRows = rocksCount * maxRockHeight

var left = &Position{x: 0, y: -1}
var right = &Position{x: 0, y: 1}
var down = &Position{x: -1, y: 0}

type Position struct {
	x     int
	y     int
	value byte
}

type Rock struct {
	name   string
	height int
	width  int
	items  []*Position
}

func (rock *Rock) canMoveToDirection(state *State, direction *Position) bool {
	if direction.y < 0 && state.x <= 0 {
		return false
	}

	if direction.y > 0 && state.x >= towerCols-rock.width {
		return false
	}

	if direction.y == 0 && state.y <= 0 {
		return false
	}

	for _, item := range rock.items {
		if item.value == ROCK && state.IsMarked(rock, item, direction) {
			return false
		}
	}

	return true
}

type State struct {
	tower  [towerRows][towerCols]bool
	x      int
	y      int
	height int
}

func (state *State) IsMarked(rock *Rock, rockItem *Position, direction *Position) bool {
	row, col := state.getPosition(rock, rockItem)

	return state.tower[row][col+direction.y]
}

func (state *State) Update(rock *Rock) {
	for _, item := range rock.items {
		if item.value == ROCK {
			row := state.mark(rock, item)
			if row > state.height {
				state.height = row
			}
		}
	}
}

func (state *State) mark(rock *Rock, rockItem *Position) int {
	row, col := state.getPosition(rock, rockItem)

	state.tower[row][col] = true

	return row
}

func (state *State) getPosition(rock *Rock, rockItem *Position) (int, int) {
	return state.y + rock.height - rockItem.x - 1, state.x + rockItem.y
}

func generateRocks() []*Rock {
	rock1 := &Rock{name: "rock1", width: 4, height: 1, items: []*Position{
		{0, 0, ROCK}, {0, 1, ROCK}, {0, 2, ROCK}, {0, 3, ROCK},
	}}
	rock2 := &Rock{name: "rock2", width: 3, height: 3, items: []*Position{
		{0, 0, EMPTY}, {0, 1, ROCK}, {0, 2, EMPTY},
		{1, 0, ROCK}, {1, 1, ROCK}, {1, 2, ROCK},
		{2, 0, EMPTY}, {2, 1, ROCK}, {2, 2, EMPTY}}}
	rock3 := &Rock{name: "rock3", width: 3, height: 3, items: []*Position{
		{0, 0, EMPTY}, {0, 1, EMPTY}, {0, 2, ROCK},
		{1, 0, EMPTY}, {1, 1, EMPTY}, {1, 2, ROCK},
		{2, 0, ROCK}, {2, 1, ROCK}, {2, 2, ROCK},
	}}
	rock4 := &Rock{name: "rock4", width: 1, height: 4, items: []*Position{
		{0, 0, ROCK},
		{1, 0, ROCK},
		{2, 0, ROCK},
		{3, 0, ROCK},
	}}
	rock5 := &Rock{name: "rock5", width: 2, height: 2, items: []*Position{
		{0, 0, ROCK}, {0, 1, ROCK},
		{1, 0, ROCK}, {1, 1, ROCK},
	}}

	return []*Rock{rock1, rock2, rock3, rock4, rock5}
}

var rockIndex = 0

func getNextRock() *Rock {
	rocks := generateRocks()

	return rocks[rockIndex%len(rocks)]
}

var jetIndex = 0
var directions = map[byte]*Position{'<': left, '>': right}

func getNextDirection() *Position {

	if jetIndex == len(input) {
		jetIndex = 0
	}
	directionChar := input[jetIndex]
	jetIndex++

	direction := directions[directionChar]

	return direction
}

func main() {
	state := &State{height: 0, x: 2, y: maxRockHeight, tower: [towerRows][towerCols]bool{}}

	for rockIndex < rocksCount {
		state.x = 2
		state.y = state.height + 4

		rock := getNextRock()

		for {
			direction := getNextDirection()

			// move left or right
			if rock.canMoveToDirection(state, direction) {
				state.x += direction.y
			}

			// move down
			state.y += down.x
			if !rock.canMoveToDirection(state, down) {
				state.y -= down.x
				break
			}
		}

		state.Update(rock)

		rockIndex++
	}

	fmt.Println(state.height)
}

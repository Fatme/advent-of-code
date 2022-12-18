package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type xyz [3]int

var lava = map[xyz]bool{}

//go:embed input.txt
var input string

var max = 0

func main() {
	input = strings.TrimSpace(input)

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		if x > max {
			max = x
		}

		if y > max {
			max = y
		}

		if z > max {
			max = z
		}

		position := xyz{x, y, z}

		lava[position] = true
	}

	solvePart1()
	solvePart2()
}

var pairs = [][3]int{
	xyz{1, 0, 0},
	xyz{-1, 0, 0},
	xyz{0, 1, 0},
	xyz{0, -1, 0},
	xyz{0, 0, 1},
	xyz{0, 0, -1},
}

func solvePart2() {
	queue := [][2]xyz{}
	queue = append(queue, [2]xyz{{-1, -1, -1}, {}})

	visited := map[xyz]bool{}
	visible := map[xyz]bool{}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		position1, position2 := item[0], item[1]

		if lava[position1] {
			visible[position2] = true
			continue
		}

		if _, ok := visited[position1]; ok {
			continue
		}

		foundInvalid := false
		for _, value := range position1 {
			if value > max+1 || value < -1 {
				foundInvalid = true
				break
			}
		}

		if foundInvalid {
			continue
		}

		visited[position1] = true

		for _, pair := range pairs {
			newPosition := xyz{
				position1[0] + pair[0],
				position1[1] + pair[1],
				position1[2] + pair[2],
			}

			queue = append(queue, [2]xyz{newPosition, position1})
		}
	}

	sum := 0

	for position, _ := range lava {
		for _, pair := range pairs {
			newPosition := xyz{
				position[0] + pair[0],
				position[1] + pair[1],
				position[2] + pair[2],
			}

			if _, ok := visible[newPosition]; ok {
				sum++
			}
		}
	}

	fmt.Println(sum)

}

func solvePart1() {
	sum := 0

	for position, _ := range lava {
		for _, pair := range pairs {
			newPosition := xyz{
				position[0] + pair[0],
				position[1] + pair[1],
				position[2] + pair[2],
			}

			if _, ok := lava[newPosition]; !ok {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = `2-4,6-8
2-3,4-5
5-7,7-9
3-7,2-8
6-6,4-6
2-6,4-8`

func processLine(line string) (int, int, int, int) {
	parts := strings.Split(line, ",")

	firstPairParts := strings.Split(parts[0], "-")
	x1, _ := strconv.Atoi(firstPairParts[0])
	x2, _ := strconv.Atoi(firstPairParts[1])

	secondPairParts := strings.Split(parts[1], "-")
	y1, _ := strconv.Atoi(secondPairParts[0])
	y2, _ := strconv.Atoi(secondPairParts[1])

	return x1, x2, y1, y2
}

func main() {
	count := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		x1, x2, y1, y2 := processLine(line)

		// Part1
		// if (x1 <= y1 && y2 <= x2) || (y1 <= x1 && x2 <= y2) {
		// 	count++
		// }

		// Part2
		if (x1 <= y1 && y1 <= x2) || (y1 <= x1 && x1 <= y2) {
			count++
		}
	}

	fmt.Println(count)
}

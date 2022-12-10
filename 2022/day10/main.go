package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	cycle := 1
	x := 1
	sum := 0

	isCompletionCycle := func(c int) bool {
		return c == 20 || c == 60 || c == 100 || c == 140 || c == 180 || c == 220
	}

	for _, line := range lines {
		cycle++

		if isCompletionCycle(cycle) {
			sum += x * cycle
		}

		parts := strings.Split(line, " ")
		command := parts[0]

		if len(parts) > 1 && command == "addx" {
			argument, _ := strconv.Atoi(parts[1])
			x += argument
			cycle++

			if isCompletionCycle(cycle) {
				sum += x * cycle
			}
		}
	}

	fmt.Println(sum)
}

package main

import (
	"fmt"
	"strings"
)

var input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func createPriorityMap() map[string]int {
	result := map[string]int{}

	priority := 1

	for r := 'a'; r <= 'z'; r++ {
		result[string(r)] = priority
		upper := strings.ToUpper(string(r))
		result[upper] = 26 + priority
		priority++
	}

	return result
}

// Part 2
func main() {
	priorities := createPriorityMap()

	sum := 0

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines)-2; i += 3 {
		firstLine := lines[i]
		secondLine := lines[i+1]
		thirdLine := lines[i+2]

		for j := 0; j < len(firstLine); j++ {
			symbol := string(firstLine[j])

			if strings.Contains(secondLine, symbol) && strings.Contains(thirdLine, symbol) {
				sum += priorities[symbol]
				break
			}
		}
	}

	fmt.Println(sum)
}

// Part 1
func main1() {
	priorities := createPriorityMap()

	sum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sizeToDivide := len(line) / 2

		part1 := line[:sizeToDivide]
		part2 := line[sizeToDivide:]

		for i := 0; i < len(part1); i++ {
			symbol := string(part1[i])

			if strings.Contains(part2, symbol) {
				sum += priorities[symbol]
				break
			}
		}
	}

	fmt.Println(sum)
}

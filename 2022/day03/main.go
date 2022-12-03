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

func main() {
	priorities := createPriorityMap()

	sum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sizeToDivide := len(line) / 2

		part1 := line[:sizeToDivide]
		part2 := line[sizeToDivide:]

		found := map[string]int{}

		for i := 0; i < len(part1); i++ {
			symbol := string(part1[i])

			if strings.Contains(part2, symbol) && found[symbol] == 0 {
				sum += priorities[symbol]
				found[symbol]++
			}
		}
	}

	fmt.Println(sum)
}

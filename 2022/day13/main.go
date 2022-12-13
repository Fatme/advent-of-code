package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	_ "embed"
)

func unmarshal(str string) interface{} {
	var parsed interface{}

	if err := json.Unmarshal([]byte(str), &parsed); err != nil {
		panic(err)
	}

	return parsed
}

// -1 -> right order and 1 -> wrong order
func compareItems(first, second []interface{}) int {
	for index, firstItem := range first {
		if len(second) <= index {
			return 1 // second runs out of items
		}

		secondItem := second[index]

		compare := func(f, s interface{}) int {
			firstNumber, firstOK := f.(float64)
			secondNumber, secondOK := s.(float64)
			if firstOK && secondOK {
				return int(firstNumber - secondNumber)
			}

			return compareItems(parseItem(f), parseItem(s))
		}

		if r := compare(firstItem, secondItem); r != 0 {
			return r
		}
	}

	if len(first) == len(second) {
		return 0
	}

	return -1
}

func parseItem(item interface{}) []interface{} {
	var result []interface{}

	switch item.(type) {
	case []interface{}, []float64:
		result = item.([]interface{})
	case float64:
		result = []interface{}{item}
	}

	return result
}

func main() {
	input, _ := os.ReadFile("input.txt")

	newLineSplitter := regexp.
		MustCompile("\r\n").
		ReplaceAllString(string(input), "\n")

	lines := regexp.
		MustCompile(`\n\s*\n`).
		Split(newLineSplitter, -1)

	sum := 0

	for index, line := range lines {
		pairs := strings.Split(line, "\n")

		first := unmarshal(pairs[0])
		second := unmarshal(pairs[1])

		result := compareItems(parseItem(first), parseItem(second))
		if result < 0 {
			sum += (index + 1)
		}
	}

	fmt.Println(sum)
}

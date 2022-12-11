package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

const rounds = 10000

// Monkey
type Monkey struct {
	items     *arraylist.List
	operation *Operation
	test      *Test
}

func (m *Monkey) calculateWorryLevel(item int64) int64 {
	operand1 := m.operation.parseOperand(m.operation.operand1, item)
	operand2 := m.operation.parseOperand(m.operation.operand2, item)

	result := int64(0)

	if m.operation.operation == "*" {
		result = operand1 * operand2
	} else if m.operation.operation == "+" {
		result = operand1 + operand2
	}

	// Part 1
	// return result / 3

	return result % (2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23)
}

func (m *Monkey) getMonkeyToThrow(worryLevel int64) int {
	if worryLevel%int64(m.test.divisibleBy) == 0 {
		return m.test.monkeyToThrowIfTrue
	}

	return m.test.monkeyToThrowIfFalse
}

// Operation
type Operation struct {
	operand1  string
	operand2  string
	operation string
}

func (o *Operation) parseOperand(operand string, item int64) int64 {
	if operand == "old" {
		return item
	}

	num, _ := strconv.Atoi(operand)
	return int64(num)
}

// Test
type Test struct {
	divisibleBy          int
	monkeyToThrowIfTrue  int
	monkeyToThrowIfFalse int
}

type Note struct {
	lines []string
}

func (n *Note) parseItems() *arraylist.List {
	itemsRaw := strings.Split(n.lines[1][18:], ", ")
	items := arraylist.New()
	for i := 0; i < len(itemsRaw); i++ {
		num, _ := strconv.Atoi(itemsRaw[i])
		items.Add(int64(num))
	}

	return items
}

func (n *Note) parseOperation() *Operation {
	parts := strings.Split(n.lines[2][13:], " ")

	return &Operation{
		operation: parts[3],
		operand1:  parts[2],
		operand2:  parts[4],
	}
}

func (n *Note) parseTest() *Test {
	parts := strings.Split(n.lines[3], " Test: divisible by ")
	divisibleBy, _ := strconv.Atoi(parts[1])

	trueTestParts := strings.Split(n.lines[4], " If true: throw to monkey ")
	monkeyToThrowIfTrue, _ := strconv.Atoi(trueTestParts[1])

	falseTestParts := strings.Split(n.lines[5], " If false: throw to monkey ")
	monkeyToThrowIfFalse, _ := strconv.Atoi(falseTestParts[1])

	return &Test{
		divisibleBy:          divisibleBy,
		monkeyToThrowIfTrue:  monkeyToThrowIfTrue,
		monkeyToThrowIfFalse: monkeyToThrowIfFalse,
	}
}

func parseInput(input string) []*Monkey {
	newLineSplitter := regexp.
		MustCompile("\r\n").
		ReplaceAllString(string(input), "\n")

	lines := regexp.
		MustCompile(`\n\s*\n`).
		Split(newLineSplitter, -1)

	monkeys := make([]*Monkey, len(lines))

	for index, line := range lines {
		note := &Note{lines: strings.Split(line, "\n")}

		items := note.parseItems()
		operation := note.parseOperation()
		test := note.parseTest()

		monkeys[index] = &Monkey{items: items, operation: operation, test: test}
	}

	return monkeys
}

func main() {
	input, _ := os.ReadFile("input.txt")

	monkeys := parseInput(string(input))

	inspectedItems := map[int]int{}

	for i := 0; i < rounds; i++ {
		for monkeyIndex, monkey := range monkeys {
			for monkey.items.Size() > 0 {
				itemRaw, _ := monkey.items.Get(0)
				item := itemRaw.(int64)

				monkey.items.Remove(0)

				worryLevel := monkey.calculateWorryLevel(item)

				monkeyToThrow := monkey.getMonkeyToThrow(worryLevel)
				monkeys[monkeyToThrow].items.Add(worryLevel)

				inspectedItems[monkeyIndex]++
			}
		}
	}

	keys := []int{}
	for key, _ := range inspectedItems {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return inspectedItems[keys[i]] > inspectedItems[keys[j]] })

	fmt.Println(inspectedItems[keys[0]] * inspectedItems[keys[1]])
}

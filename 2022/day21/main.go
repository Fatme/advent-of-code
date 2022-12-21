package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	name      string
	value     int
	operation string
	operator  string
	monkey1   string
	monkey2   string
}

func (monkey *Monkey) setValues() {
	operator := ""

	if strings.Contains(monkey.operation, "+") {
		operator = "+"
	} else if strings.Contains(monkey.operation, "-") {
		operator = "-"
	} else if strings.Contains(monkey.operation, "*") {
		operator = "*"
	} else if strings.Contains(monkey.operation, "/") {
		operator = "/"
	}

	parts := strings.Split(monkey.operation, fmt.Sprintf(" %s ", operator))

	monkey.monkey1 = parts[0]
	monkey.monkey2 = parts[1]
	monkey.operator = operator
}

func (monkey *Monkey) process() {
	if !isNumber(monkey.operation) {
		monkey1 := monkeysMap[monkey.monkey1]
		monkey2 := monkeysMap[monkey.monkey2]

		if monkey1.value != 0 && monkey2.value != 0 {
			result := 0
			if monkey.operator == "+" {
				result = monkey1.value + monkey2.value
			} else if monkey.operator == "-" {
				result = monkey1.value - monkey2.value
			} else if monkey.operator == "*" {
				result = monkey1.value * monkey2.value
			} else if monkey.operator == "/" {
				result = monkey1.value / monkey2.value
			}

			monkey.value = result
			monkey.operation = strconv.Itoa(result)
			monkey.monkey1 = ""
			monkey.monkey2 = ""
		}
	}
}

var monkeysMap = map[string]*Monkey{}

const rootMonkey = "root"

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		monkeyName := parts[0]
		monkey := &Monkey{name: monkeyName}

		if num, err := strconv.Atoi(parts[1]); err == nil {
			monkey.operation = strconv.Itoa(num)
			monkey.value = num
		} else {
			monkey.operation = parts[1]
			monkey.setValues()
		}

		monkeysMap[monkeyName] = monkey
	}

	solvePart1()
}

func solvePart1() {
	for !isNumber(monkeysMap[rootMonkey].operation) {
		for _, monkey := range monkeysMap {
			monkey.process()
		}
	}

	fmt.Println(monkeysMap[rootMonkey].value)
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

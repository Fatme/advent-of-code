package main

import (
	"fmt"
	"math"
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

func (monkey *Monkey) getValue() float64 {
	if isNumber(monkey.operation) {
		return float64(monkey.value)
	}

	if monkey.operator == "+" {
		return monkeysMap[monkey.monkey1].getValue() + monkeysMap[monkey.monkey2].getValue()
	} else if monkey.operator == "-" {
		return monkeysMap[monkey.monkey1].getValue() - monkeysMap[monkey.monkey2].getValue()
	} else if monkey.operator == "*" {
		return monkeysMap[monkey.monkey1].getValue() * monkeysMap[monkey.monkey2].getValue()
	} else if monkey.operator == "/" {
		return monkeysMap[monkey.monkey1].getValue() / monkeysMap[monkey.monkey2].getValue()
	}

	panic("unexpected")
}

var monkeysMap = map[string]*Monkey{}

const rootMonkey = "root"
const myMonkey = "humn"

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
			data := strings.Split(parts[1], " ")

			monkey.monkey1 = data[0]
			monkey.monkey2 = data[2]
			monkey.operator = data[1]
			monkey.operation = parts[1]
		}

		monkeysMap[monkeyName] = monkey
	}

	// solvePart1()
	solvePart2()
}

func solvePart1() {
	for !isNumber(monkeysMap[rootMonkey].operation) {
		for _, monkey := range monkeysMap {
			monkey.process()
		}
	}

	fmt.Println(monkeysMap[rootMonkey].value)
}

func solvePart2() {
	var min = 0
	var max = math.MaxInt

	monkey := monkeysMap[rootMonkey]

	for min < max {
		mid := (min + max) / 2

		monkeysMap[myMonkey].value = mid
		monkeysMap[myMonkey].operation = strconv.Itoa(mid)

		value1 := monkeysMap[monkey.monkey1].getValue()
		value2 := monkeysMap[monkey.monkey2].getValue()

		if value1 == value2 {
			fmt.Println(mid)
			return
		}

		if value1 > value2 {
			min = mid + 1
		} else {
			max = mid
		}
	}
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

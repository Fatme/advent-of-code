package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

const ROWS = 10 // 4
const COLS = 9  // 3

const INPUT_FILE = "input.txt"

var data [COLS]*arraylist.List

func main() {
	instructionsLine := 0

	input, _ := os.ReadFile(INPUT_FILE)

	lines := strings.Split(string(input), "\n")

	for index, line := range lines {
		if index == ROWS-2 {
			instructionsLine = index + 1
			break
		}

		col := 0

		for j := 0; j < len(line); j += 4 {
			item := string(line[j+1])
			// Ignore empty cells
			if item == " " {
				col++
				continue
			}

			if data[col] == nil {
				data[col] = arraylist.New()
			}

			data[col].Insert(0, item)
			col++
		}
	}

	for index := instructionsLine + 1; index < len(lines); index++ {
		line := lines[index]

		parts := strings.Split(line, " ")

		countToMove, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		fromList := data[from-1]
		toList := data[to-1]

		for i := countToMove - 1; i >= 0; i-- {
			fromIndex := fromList.Size() - 1
			item, _ := fromList.Get(fromIndex)
			fromList.Remove(fromIndex)
			toList.Add(item)
		}
	}

	result := ""

	for i := 0; i < len(data); i++ {
		item, _ := data[i].Get(data[i].Size() - 1)
		result += item.(string)
	}

	fmt.Println(result)
}

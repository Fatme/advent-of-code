package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func main() {
	lines := strings.Split(input, "\n")

	sum := 0

	max1 := 0
	max2 := 0
	max3 := 0

	for i := 0; i < len(lines); i++ {
		num, _ := strconv.Atoi(lines[i])
		sum += num

		if lines[i] == "" || i == len(lines)-1 {
			if sum > max1 {
				max3 = max2
				max2 = max1
				max1 = sum
			} else if sum > max2 && sum != max1 {
				max3 = max2
				max2 = sum
			} else if sum > max3 && sum != max2 {
				max3 = sum
			}

			sum = 0
		}
	}

	fmt.Println(max1 + max2 + max3)
}

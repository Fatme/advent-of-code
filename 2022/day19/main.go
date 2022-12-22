package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const ORE = "ore"
const CLAY = "clay"
const OBSIDIAN = "obsidian"
const GEODE = "geode"

const MAX_ALLOWED_QUEUE_SIZE = 1000000

type robots [4]int
type prices [4]int

type Blueprint struct {
	index           int
	mapRobotToPrice map[string]map[string]int
}

func canBuyOreRobot(ore int, priceMap map[string]int) bool {
	return ore >= priceMap[ORE]
}

func canBuyClayRobot(ore int, priceMap map[string]int) bool {
	return ore >= priceMap[ORE]
}

func canBuyObsidianRobot(ore, clay int, priceMap map[string]int) bool {
	return ore >= priceMap[ORE] && clay >= priceMap[CLAY]
}

func canBuyGeodeRobot(ore, obsidian int, priceMap map[string]int) bool {
	return ore >= priceMap[ORE] && obsidian >= priceMap[OBSIDIAN]
}

func (b *Blueprint) process(minutesToPlay int) int {
	max := 0

	visited := map[State]bool{}

	queue := []State{{robots: robots{1, 0, 0, 0}, prices: prices{0, 0, 0, 0}}}

	for len(queue) > 0 {

		if len(queue) > MAX_ALLOWED_QUEUE_SIZE {
			// Order the queue by priority
			sort.Slice(queue, func(i, j int) bool {
				sum1 := queue[i].robots[3]*10000 + queue[i].robots[2]*1000 + queue[i].robots[1]*100 + queue[i].robots[0]*10
				sum2 := queue[j].robots[3]*10000 + queue[j].robots[2]*1000 + queue[j].robots[1]*100 + queue[j].robots[0]*10
				return sum1 > sum2
			})
			queue = queue[0:MAX_ALLOWED_QUEUE_SIZE]
			continue
		}

		state := queue[0]
		queue = queue[1:]

		if state.minutes == minutesToPlay {
			if state.prices[3] > max {
				max = state.prices[3]
			}

			continue
		}

		if visited[state] {
			continue
		}

		visited[state] = true

		if canBuyGeodeRobot(state.prices[0], state.prices[2], b.mapRobotToPrice[GEODE]) {
			queue = append(queue, state.createFromState(GEODE, b.mapRobotToPrice[GEODE]))
		}

		if canBuyObsidianRobot(state.prices[0], state.prices[1], b.mapRobotToPrice[OBSIDIAN]) {
			queue = append(queue, state.createFromState(OBSIDIAN, b.mapRobotToPrice[OBSIDIAN]))
		}

		if canBuyClayRobot(state.prices[0], b.mapRobotToPrice[CLAY]) {
			queue = append(queue, state.createFromState(CLAY, b.mapRobotToPrice[CLAY]))
		}

		if canBuyOreRobot(state.prices[0], b.mapRobotToPrice[ORE]) {
			queue = append(queue, state.createFromState(ORE, b.mapRobotToPrice[ORE]))
		}

		queue = append(queue, state.createFromState("", map[string]int{}))
	}

	return max
}

type State struct {
	robots  robots
	prices  prices
	minutes int
}

func (s *State) createFromState(robot string, priceMap map[string]int) State {
	newPrices := prices{s.prices[0] + s.robots[0], s.prices[1] + s.robots[1], s.prices[2] + s.robots[2], s.prices[3] + s.robots[3]}

	result := State{
		robots:  s.robots,
		prices:  newPrices,
		minutes: s.minutes + 1}

	if robot == ORE {
		result.prices[0] -= priceMap[ORE]

		result.robots[0]++
	} else if robot == CLAY {
		result.prices[0] -= priceMap[ORE]

		result.robots[1]++
	} else if robot == OBSIDIAN {
		result.prices[0] -= priceMap[ORE]
		result.prices[1] -= priceMap[CLAY]

		result.robots[2]++
	} else if robot == GEODE {
		result.prices[0] -= priceMap[ORE]
		result.prices[2] -= priceMap[OBSIDIAN]

		result.robots[3]++
	}

	return result
}

func parseInput(lines []string) []*Blueprint {
	result := []*Blueprint{}

	regex := regexp.MustCompile(`-?\d+`)

	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		numbers := make([]int, len(matches))
		for key, value := range matches {
			numbers[key], _ = strconv.Atoi(value)
		}

		blueprint := &Blueprint{
			index: numbers[0],
			mapRobotToPrice: map[string]map[string]int{
				ORE:      {ORE: numbers[1]},
				CLAY:     {ORE: numbers[2]},
				OBSIDIAN: {ORE: numbers[3], CLAY: numbers[4]},
				GEODE:    {ORE: numbers[5], OBSIDIAN: numbers[6]},
			},
		}

		result = append(result, blueprint)
	}

	return result
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	blueprints := parseInput(lines)

	solvePart1(blueprints)
	// solvePart2(blueprints)
}

func solvePart1(blueprints []*Blueprint) {
	sum := 0

	for _, blueprint := range blueprints {
		result := blueprint.process(24)
		sum += blueprint.index * result
	}

	fmt.Println(sum)

}

func solvePart2(blueprints []*Blueprint) {
	result := 1

	for _, blueprint := range blueprints[:3] {
		result *= blueprint.process(32)
	}

	fmt.Println("part2 result", result)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[i] = num
		}

		isIncreasing := true
		isSafe := true

		for i := 0; i < len(levels)-1; i++ {
			difference := levels[i] - levels[i+1]

			if difference == 0 {
				isSafe = false
				break
			}

			if difference < -3 || difference > 3 {
				isSafe = false
				break
			}

			if i == 0 && difference > 0 {
				isIncreasing = false
			} else if isIncreasing && difference > 0 || !isIncreasing && difference < 0 {
				isSafe = false
				break
			}
		}
		if isSafe {
			output++
		}
	}

	fmt.Println("Day 2 Part 1 Output is", output)
}

func day2_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[i] = num
		}

		isSafe := safe(levels)

		if !isSafe {
			for i := 0; i < len(levels); i++ {
				modifiedLevels := []int{}
				modifiedLevels = append(modifiedLevels, levels[:i]...)
				modifiedLevels = append(modifiedLevels, levels[i+1:]...)
				isSafe = safe(modifiedLevels)
				if isSafe {
					break
				}
			}
		}
		if isSafe {
			output++
		}
	}
	fmt.Println("Day 2 Part 2 Output is", output)
}

func safe(levels []int) bool {
	isIncreasing := true
	isSafe := true

	for i := 0; i < len(levels)-1; i++ {
		difference := levels[i] - levels[i+1]

		if difference == 0 {
			isSafe = false
			break
		}

		if difference < -3 || difference > 3 {
			isSafe = false
			break
		}

		if i == 0 && difference > 0 {
			isIncreasing = false
		} else if isIncreasing && difference > 0 || !isIncreasing && difference < 0 {
			isSafe = false
			break
		}
	}
	return isSafe
}

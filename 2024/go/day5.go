package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day5_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	rules := make(map[int][]int)
	var updates [][]int

	isRule := true

	for _, line := range lines {
		if line == "" {
			isRule = false
			continue
		}
		if isRule {
			parts := strings.Split(line, "|")
			key, _ := strconv.Atoi(parts[0])
			value, _ := strconv.Atoi(parts[1])
			rules[key] = append(rules[key], value)
		} else {
			var currentUpdate []int
			parts := strings.Split(line, ",")
			for _, part := range parts {
				num1, _ := strconv.Atoi(part)
				currentUpdate = append(currentUpdate, num1)
			}

			updates = append(updates, currentUpdate)
		}

	}

	for _, currentUpdate := range updates {
		isValid := true

		for i := len(currentUpdate) - 1; i > 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, rule := range rules[currentNum] {
					if num == rule {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}

		}
		if isValid {
			middle := currentUpdate[len(currentUpdate)/2]
			output += middle
		}
	}

	fmt.Println("Day 5 Part 1 Output is", output)
}

func day5_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	rules := make(map[int][]int)
	var updates [][]int

	isRule := true

	for _, line := range lines {
		if line == "" {
			isRule = false
			continue
		}
		if isRule {
			parts := strings.Split(line, "|")
			key, _ := strconv.Atoi(parts[0])
			value, _ := strconv.Atoi(parts[1])
			rules[key] = append(rules[key], value)
		} else {
			var currentUpdate []int
			parts := strings.Split(line, ",")
			for _, part := range parts {
				num1, _ := strconv.Atoi(part)
				currentUpdate = append(currentUpdate, num1)
			}

			updates = append(updates, currentUpdate)
		}

	}

	for _, currentUpdate := range updates {
		isValid := true

		for i := len(currentUpdate) - 1; i > 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, rule := range rules[currentNum] {
					if num == rule {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}

		}
		if !isValid {
			ordered := []int{}
			remaining := map[int]bool{}
			dependencies := make(map[int][]int)

			for _, num := range currentUpdate {
				remaining[num] = true

				for _, dep := range rules[num] {
					dependencies[dep] = append(dependencies[dep], num)
				}
			}

			for len(remaining) > 0 {
				for num := range remaining {
					if len(dependencies[num]) == 0 {
						ordered = append(ordered, num)
						delete(remaining, num)

						for key, val := range dependencies {
							newList := []int{}
							for _, n := range val {
								if n != num {
									newList = append(newList, n)
								}
							}
							dependencies[key] = newList
						}
					}
				}
			}
			middle := ordered[len(ordered)/2]
			output += middle
		}
	}

	fmt.Println("Day 5 Part 2 Output is", output)
}

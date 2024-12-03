package day3

import (
	"awesomeProject/utils"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func searchNext(input, token string, position int) (int, error) {
	for i := position; i < len(input); i++ {
		if input[i] == token[0] {
			tokenFound := true
			for j := 1; j < len(token); j++ {
				if i+j < len(input) {
					if input[i+j] != token[j] {
						tokenFound = false
						break
					}
				}
			}

			if tokenFound {
				return i + len(token), nil
			}
		}
	}

	return 0, errors.New(fmt.Sprintf("Could not find token %s after position %d", token, position))
}

func Part1() int {
	lines := utils.ReadInput("day3")
	input := strings.Join(lines, "")

	re, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")
	matches := re.FindAllString(input, -1)

	return calculateMuls(matches)
}

func calculateMuls(matches []string) int {
	var total int64

	mulsEnabled := true

	for _, match := range matches {
		if match == "don't()" {
			mulsEnabled = false
			continue
		}

		if match == "do()" {
			mulsEnabled = true
			continue
		}

		if !mulsEnabled {
			continue
		}

		noMul := strings.ReplaceAll(match, "mul(", "")
		cleanMul := strings.ReplaceAll(noMul, ")", "")
		parts := strings.Split(cleanMul, ",")

		num1, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			log.Fatalf("Failed to parse input as number: %s", parts[0])
		}
		num2, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			log.Fatalf("Failed to parse input as number: %s", parts[1])
		}

		total += num1 * num2
	}

	return int(total)
}

func Part2() int {
	lines := utils.ReadInput("day3")
	input := strings.Join(lines, "")

	re, _ := regexp.Compile("mul\\(\\d+,\\d+\\)|don't\\(\\)|do\\(\\)")
	matches := re.FindAllString(input, -1)

	return calculateMuls(matches)
}

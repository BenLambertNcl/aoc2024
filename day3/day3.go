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

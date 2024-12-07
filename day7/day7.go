package day7

import (
	"awesomeProject/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Part1() int {
	lines := utils.ReadInput("day7")
	return parseInput(lines, []Operator{MULT, ADD})
}

func Part2() int {
	lines := utils.ReadInput("day7")
	return parseInput(lines, []Operator{MULT, ADD, CONCAT})
}

func parseInput(lines []string, validOperators []Operator) int {
	var totalAnswers int64 = 0

	for _, line := range lines {
		parts := strings.Split(line, ":")
		answer, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		sumParts := strings.Split(parts[1], " ")
		var intParts []int64
		for _, s := range sumParts {
			if s == "" {
				continue
			}
			n, err := strconv.ParseInt(s, 10, 0)
			if err != nil {
				log.Fatal(err)
			}
			intParts = append(intParts, n)
		}

		if isValidInput(intParts, answer, validOperators) {
			totalAnswers += answer
		}
	}

	return int(totalAnswers)
}

type Operator int

const (
	MULT Operator = iota
	ADD
	CONCAT
)

func isValidInput(sumParts []int64, answer int64, validOperators []Operator) bool {
	operatorsPerLine := len(sumParts) - 1
	combinationsNeeded := powInt(len(validOperators), operatorsPerLine)
	var allCombinations [][]Operator

	for i := range combinationsNeeded {
		binaryString := fmt.Sprintf("%0*s", operatorsPerLine, strconv.FormatInt(int64(i), len(validOperators)))
		var combinations []Operator
		for _, char := range binaryString {
			num, err := strconv.ParseInt(string(char), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			combinations = append(combinations, validOperators[int(num)])
		}
		allCombinations = append(allCombinations, combinations)
	}

	for _, combination := range allCombinations {
		var total int64 = 0
		prev := sumParts[0]
		for i := 1; i < len(sumParts); i++ {
			current := sumParts[i]
			operator := combination[i-1]
			total = doSum(prev, operator, current)
			prev = total
		}
		if total == answer {
			return true
		}
	}

	return false
}

func doSum(left int64, operator Operator, right int64) int64 {
	switch operator {
	case MULT:
		return left * right
	case ADD:
		return left + right
	case CONCAT:
		num, err := strconv.ParseInt(
			strings.Join([]string{
				strconv.FormatInt(left, 10),
				strconv.FormatInt(right, 10),
			},
				"",
			),
			10,
			0,
		)
		if err != nil {
			log.Fatal(err)
		}
		return num
	default:
		log.Fatalf("Unknown operator: %v", operator)
		return -1
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

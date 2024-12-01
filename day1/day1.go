package day1

import (
	"awesomeProject/utils"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1() int {
	lines := utils.ReadInput("day1/input.txt")
	leftList, rightList := createLists(lines)

	sort.Ints(leftList)
	sort.Ints(rightList)

	var distances []int

	for i := 0; i < len(leftList); i++ {
		distance := math.Abs(float64(leftList[i] - rightList[i]))
		distances = append(distances, int(distance))
	}

	total := 0

	for _, distance := range distances {
		total += distance
	}

	return total
}

func Part2() int {
	lines := utils.ReadInput("day1/input.txt")
	leftList, rightList := createLists(lines)

	total := 0

	for _, leftNum := range leftList {
		occurrences := 0
		for _, rightNum := range rightList {
			if leftNum == rightNum {
				occurrences++
			}
		}
		total += leftNum * occurrences
	}

	return total
}

func createLists(lines []string) ([]int, []int) {
	var leftList []int
	var rightList []int

	for _, line := range lines {
		parts := strings.Split(line, "   ")

		leftNum, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		rightNum, err := strconv.ParseInt(parts[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		leftList = append(leftList, int(leftNum))
		rightList = append(rightList, int(rightNum))
	}
	return leftList, rightList
}

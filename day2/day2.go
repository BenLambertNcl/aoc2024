package day2

import (
	"awesomeProject/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func Part1() int {
	reports := utils.ReadInput("day2")

	safeReports := 0
	for _, report := range reports {
		badLevel := findBadLevel(report)
		if badLevel == -1 {
			safeReports++
			continue
		}
	}

	return safeReports
}

/*
*
Returns -1 if there are no bad levels
*/
func findBadLevel(report string) int {
	reportParts := strings.Split(report, " ")
	var previousNum int64
	isAscending := true
	for i, reportPart := range reportParts {
		reportPartNum, err := strconv.ParseInt(reportPart, 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			previousNum = reportPartNum
			continue
		}

		if i == 1 {
			isAscending = reportPartNum > previousNum
		}

		if isAscending && reportPartNum <= previousNum {
			return i
		}

		if !isAscending && reportPartNum >= previousNum {
			return i
		}
		difference := int(math.Abs(float64(reportPartNum - previousNum)))
		if difference > 3 {
			return i
		}

		previousNum = reportPartNum
	}
	return -1
}

func Part2() int {
	reports := utils.ReadInput("day2")

	safeReports := 0

outer:
	for _, report := range reports {
		badLevel := findBadLevel(report)
		if badLevel == -1 {
			safeReports++
			continue
		}

		parts := strings.Split(report, " ")
		unmodifiedParts := strings.Split(report, " ")

		for i, _ := range parts {
			var newParts []string

			for j, val := range unmodifiedParts {
				if i == j {
					continue
				}
				newParts = append(newParts, val)
			}

			badLevel = findBadLevel(strings.Join(newParts, " "))
			if badLevel == -1 {
				safeReports++
				continue outer
			}
		}
	}

	return safeReports
}

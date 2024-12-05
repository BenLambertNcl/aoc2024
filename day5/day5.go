package day5

import (
	"awesomeProject/utils"
	"errors"
	"log"
	"strconv"
	"strings"
)

type Rule struct {
	containsBefore string
	number         string
}

func Part1() int {
	lines := utils.ReadInput("day5")
	updates, rules := parsePuzzleInput(lines)

	var total int64 = 0
	for _, update := range updates {
		middle, err := getMiddleOfUpdate(update, rules)
		if err != nil {
			continue
		}
		total += middle
	}

	return int(total)
}

func Part2() int {
	lines := utils.ReadInput("day5")
	updates, rules := parsePuzzleInput(lines)

	var incorrectUpdates []string
	for _, update := range updates {
		_, err := getMiddleOfUpdate(update, rules)
		if err != nil {
			incorrectUpdates = append(incorrectUpdates, update)
			continue
		}
	}

	reordered := reorderIncorrectUpdates(incorrectUpdates, rules)

	var total int64 = 0
	for _, update := range reordered {
		middle, err := getMiddleOfUpdate(update, rules)
		if err != nil {
			continue
		}
		total += middle
	}

	return int(total)
}

func getMiddleOfUpdate(update string, rules []Rule) (int64, error) {
	for _, rule := range rules {
		location := strings.Index(update, rule.number)
		if location != -1 {
			locationPrevious := strings.Index(update, rule.containsBefore)
			if locationPrevious > location {
				return 0, errors.New("not a valid update")
			}
		}
	}
	nums := strings.Split(update, ",")
	middle, err := strconv.ParseInt(nums[len(nums)/2], 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	return middle, nil
}

func parsePuzzleInput(lines []string) ([]string, []Rule) {
	var rules []Rule
	var updates []string

	endOfRules := false

	for _, line := range lines {
		if line == "" {
			endOfRules = true
			continue
		}

		if !endOfRules {
			parts := strings.Split(line, "|")
			rule := Rule{
				containsBefore: parts[0],
				number:         parts[1],
			}
			rules = append(rules, rule)
		} else {
			updates = append(updates, line)
		}
	}

	return updates, rules
}

func reorderIncorrectUpdates(updates []string, rules []Rule) []string {
	var newUpdates []string
	for _, update := range updates {
		reordered := reorderUpdate(update, rules)
		newUpdates = append(newUpdates, reordered)
	}
	return newUpdates
}

func reorderUpdate(update string, rules []Rule) string {
	for _, rule := range rules {
		location := strings.Index(update, rule.number)
		if location != -1 {
			locationPrevious := strings.Index(update, rule.containsBefore)
			if locationPrevious > location {
				parts := strings.Split(update, ",")
				var newUpdateParts []string
				for i := 0; i < len(parts); i++ {

					if parts[i] == rule.containsBefore {
						continue
					}

					if parts[i] == rule.number {
						newUpdateParts = append(newUpdateParts, rule.containsBefore, parts[i])
						continue
					}

					// else
					newUpdateParts = append(newUpdateParts, parts[i])
				}

				// Recursive call
				return reorderUpdate(strings.Join(newUpdateParts, ","), rules)
			}
		}
	}

	// Did not need to modify, return the input
	return update
}

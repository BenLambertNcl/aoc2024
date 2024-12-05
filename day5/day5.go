package day5

import (
	"awesomeProject/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Rule struct {
	containsBefore int64
	number         int64
}

func Part1() int {
	lines := utils.ReadInput("day5")
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
			containsBefore, err := strconv.ParseInt(parts[0], 10, 0)
			if err != nil {
				log.Fatal(err)
			}
			number, err := strconv.ParseInt(parts[1], 10, 0)
			if err != nil {
				log.Fatal(err)
			}

			rule := Rule{
				containsBefore: containsBefore,
				number:         number,
			}
			rules = append(rules, rule)
		} else {
			updates = append(updates, line)
		}
	}

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

func getMiddleOfUpdate(update string, rules []Rule) (int64, error) {
	for _, rule := range rules {
		location := strings.Index(update, strconv.FormatInt(rule.number, 10))
		if location != -1 {
			locationPrevious := strings.Index(update, strconv.FormatInt(rule.containsBefore, 10))
			if locationPrevious > location {
				return 0, errors.New("Not a valid update")
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

func Part2() int {
	lines := utils.ReadInput("day5")

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
			containsBefore, err := strconv.ParseInt(parts[0], 10, 0)
			if err != nil {
				log.Fatal(err)
			}
			number, err := strconv.ParseInt(parts[1], 10, 0)
			if err != nil {
				log.Fatal(err)
			}

			rule := Rule{
				containsBefore: containsBefore,
				number:         number,
			}
			rules = append(rules, rule)
		} else {
			updates = append(updates, line)
		}
	}

	var incorrectUpdates []string
	for _, update := range updates {
		_, err := getMiddleOfUpdate(update, rules)
		if err != nil {
			incorrectUpdates = append(incorrectUpdates, update)
			continue
		}
	}

	// Reorder
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

func reorderIncorrectUpdates(updates []string, rules []Rule) []string {
	var newUpdates []string
	for _, update := range updates {
		reordered := reorderUpdate(update, rules)
		fmt.Printf("Update %s has become %s\n", update, reordered)
		newUpdates = append(newUpdates, reordered)
	}
	return newUpdates
}

func reorderUpdate(update string, rules []Rule) string {
	for _, rule := range rules {
		location := strings.Index(update, strconv.FormatInt(rule.number, 10))
		if location != -1 {
			locationPrevious := strings.Index(update, strconv.FormatInt(rule.containsBefore, 10))
			if locationPrevious > location {
				parts := strings.Split(update, ",")
				var newUpdateParts []string
				for i := 0; i < len(parts); i++ {
					partNum, err := strconv.ParseInt(parts[i], 10, 0)
					if err != nil {
						log.Fatal(err)
					}

					if partNum == rule.containsBefore {
						continue
					}

					if partNum == rule.number {
						before := strconv.FormatInt(rule.containsBefore, 10)
						newUpdateParts = append(newUpdateParts, before, parts[i])
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

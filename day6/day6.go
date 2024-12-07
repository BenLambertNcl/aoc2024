package day6

import (
	"awesomeProject/utils"
	"errors"
	"fmt"
	"log"
	"slices"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func Part1() int {
	lines := utils.ReadInput("day6")

	var grid [][]rune
	var guardX int
	var guardY int

	for i, line := range lines {
		grid = append(grid, []rune{})
		for j, char := range line {
			grid[i] = append(grid[i], char)

			if string(char) == "^" {
				guardX, guardY = j, i
			}
		}
	}

	stepCount, _ := walkGrid(guardX, guardY, &grid)

	return stepCount
}

func Part2() int {
	lines := utils.ReadInput("day6")

	var grid [][]rune
	var guardX int
	var guardY int

	for i, line := range lines {
		grid = append(grid, []rune{})
		for j, char := range line {
			grid[i] = append(grid[i], char)

			if string(char) == "^" {
				guardX, guardY = j, i
			}
		}
	}

	loopCounter := 0

	for i, row := range grid {
		for j, char := range row {
			if string(char) == "^" || string(char) == "#" {
				continue
			}

			gridCopy := copyGrid(grid)
			gridCopy[i][j] = '#'

			_, err := walkGrid(guardX, guardY, &gridCopy)
			if err != nil {
				fmt.Println(err)
				loopCounter++
			}
		}
	}

	return loopCounter
}

type Step struct {
	x         int
	y         int
	direction Direction
}

func walkGrid(startX, startY int, grid *[][]rune) (int, error) {
	guardX := startX
	guardY := startY

	direction := UP
	stepCount := 1
	var previousSteps []Step

	for {
		// LOOP DETECTED!
		currentStep := Step{
			guardX, guardY, direction,
		}
		if stepCount != 1 && slices.Contains(previousSteps, currentStep) {
			return -1, errors.New("loop detected")
		}
		previousSteps = append(previousSteps, Step{guardX, guardY, direction})

		newX, newY := moveGuard(guardX, guardY, direction)

		if isOutOfBounds(*grid, newX, newY) {
			break
		}

		gridChar := (*grid)[newY][newX]

		if gridChar == '#' {
			direction = changeDirection(direction)
			// Need to continue now, so the guard doesn't actually move to the spot with the object
			continue
		}

		if gridChar != 'X' && gridChar != '^' {
			stepCount++
			(*grid)[newY][newX] = 'X'
		}

		guardX = newX
		guardY = newY
	}
	return stepCount, nil
}

func moveGuard(x, y int, direction Direction) (int, int) {
	switch direction {
	case UP:
		return x, y - 1
	case DOWN:
		return x, y + 1
	case LEFT:
		return x - 1, y
	case RIGHT:
		return x + 1, y
	default:
		log.Fatal("Unknown direction")
		return -1, -1
	}
}

func changeDirection(current Direction) Direction {
	switch current {
	case UP:
		return RIGHT
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	default:
		log.Fatal("Unknown direction")
		return -1
	}
}

func isOutOfBounds(grid [][]rune, x, y int) bool {
	gridY := len(grid)
	gridX := len(grid[0])
	return x < 0 || y < 0 || x >= gridX || y >= gridY
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]rune, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}

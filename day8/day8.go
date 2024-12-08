package day8

import (
	"awesomeProject/utils"
	"slices"
	"strings"
)

func Part1() int {
	lines := utils.ReadInput("day8")
	var grid [][]string

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	antennaPositions := make(map[string][]Position)
	allPositions := make([]Position, 0)

	for i, line := range grid {
		for j, char := range line {
			if string(char) != "#" && string(char) != "." {
				antennaPositions[char] = append(antennaPositions[char], Position{j, i})
				allPositions = append(allPositions, Position{j, i})
			}
		}
	}

	var allAntinodes []Position

	for _, val := range antennaPositions {
		antinodes := findAntinodes(val)

		for _, antinode := range antinodes {
			// If it's the same position as an existing antenna, the antenna wins
			if slices.Contains(allPositions, antinode) || slices.Contains(allAntinodes, antinode) {
				continue
			}

			if antinode.X >= 0 && antinode.X < len(grid[0]) && antinode.Y >= 0 && antinode.Y < len(grid) {
				allAntinodes = append(allAntinodes, antinode)
			}
		}
	}

	for _, node := range allAntinodes {
		grid[node.Y][node.X] = "#"
	}

	return len(allAntinodes)
}

func Part2() int {
	lines := utils.ReadInput("day8")
	var grid [][]string

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	antennaPositions := make(map[string][]Position)
	allPositions := make([]Position, 0)

	for i, line := range grid {
		for j, char := range line {
			if string(char) != "#" && string(char) != "." {
				antennaPositions[char] = append(antennaPositions[char], Position{j, i})
				allPositions = append(allPositions, Position{j, i})
			}
		}
	}

	var allAntinodes []Position

	for _, val := range antennaPositions {
		antinodes := findAntinodes2(val, len(grid[0]), len(grid))

		for _, antinode := range antinodes {
			if slices.Contains(allPositions, antinode) || slices.Contains(allAntinodes, antinode) {
				continue
			}

			allAntinodes = append(allAntinodes, antinode)
		}

		if len(val) > 1 {
			allAntinodes = append(allAntinodes, val...)
		}
	}

	for _, node := range allAntinodes {
		grid[node.Y][node.X] = "#"
	}

	return len(allAntinodes)
}

type Position struct {
	X int
	Y int
}

func findAntinodes(antennas []Position) []Position {
	var antinodes []Position

	for _, position := range antennas {
		for _, otherPosition := range antennas {
			if position == otherPosition {
				continue
			}

			diffX := position.X - otherPosition.X
			diffY := position.Y - otherPosition.Y

			pos1 := Position{
				X: position.X - diffX,
				Y: position.Y - diffY,
			}

			if pos1 != otherPosition {
				antinodes = append(antinodes, pos1)
			}

			pos2 := Position{
				X: position.X + diffX,
				Y: position.Y + diffY,
			}

			if pos2 != otherPosition {
				antinodes = append(antinodes, pos2)
			}
		}
	}

	return antinodes
}

func findAntinodes2(antennas []Position, maxX, maxY int) []Position {
	var antinodes []Position

	for _, position := range antennas {
		for _, otherPosition := range antennas {
			if position == otherPosition {
				continue
			}

			diffX := position.X - otherPosition.X
			diffY := position.Y - otherPosition.Y

			scaledDiffX := diffX
			scaledDiffY := diffY

			foundAnother := false

			for {
				pos1 := Position{
					X: position.X - scaledDiffX,
					Y: position.Y - scaledDiffY,
				}

				if pos1 != otherPosition && pos1.X >= 0 && pos1.X < maxX && pos1.Y >= 0 && pos1.Y < maxY {
					antinodes = append(antinodes, pos1)
					foundAnother = true
				}

				pos2 := Position{
					X: position.X + scaledDiffX,
					Y: position.Y + scaledDiffY,
				}

				if pos2 != otherPosition && pos2.X >= 0 && pos2.X < maxX && pos2.Y >= 0 && pos2.Y < maxY {
					antinodes = append(antinodes, pos2)
					foundAnother = true
				}

				if foundAnother {
					scaledDiffX += diffX
					scaledDiffY += diffY
					foundAnother = false
				} else {
					break
				}
			}
		}
	}

	return antinodes
}

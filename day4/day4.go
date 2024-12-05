package day4

import (
	"awesomeProject/utils"
	"fmt"
)

const WordToFind = "XMAS"

func Part1() int {
	lines := utils.ReadInput("day4")

	lenOfWord := len(WordToFind)

	matchCount := 0

	for i, line := range lines {
		for j, char := range line {
			if char != int32(WordToFind[0]) {
				continue
			}

			fmt.Printf("X at: (%d, %d)\n", j, i)

			//vertical
			if i+(lenOfWord-1) < len(lines) {
				// Down
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(lines[i+wordParts][j])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found vertical (t2b) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if i-(lenOfWord-1) >= 0 {
				// Up
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(lines[i-wordParts][j])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found vertical (b2t) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			//horizontal
			if j+(lenOfWord-1) < len(line) {
				// left to right
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(line[j+wordParts])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found horizontal (l2r) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if j-(lenOfWord-1) >= 0 {
				// right to left
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(line[j-wordParts])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found horizontal (r2l) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			//diagonal
			if i+(lenOfWord-1) < len(lines) && j+(lenOfWord-1) < len(line) {
				// top to bottom, left to right
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(lines[i+wordParts][j+wordParts])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found diagonal (t2b, l2r) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if i-(lenOfWord-1) >= 0 && j+(lenOfWord-1) < len(line) {
				// bottom to top, left to right
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(lines[i-wordParts][j+wordParts])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found diagonal (b2t,l2r) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if i+(lenOfWord-1) < len(lines) && j-(lenOfWord-1) >= 0 {
				// top to bottom, right to left
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(lines[i+wordParts][j-wordParts])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found diagonal (t2b,r2l) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if i-(lenOfWord-1) >= 0 && j-(lenOfWord-1) >= 0 {
				// bottom to top, right to left
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(lines[i-wordParts][j-wordParts])
				}

				if foundWord == WordToFind {
					fmt.Printf("Found diagonal (b2t,r2l) at (%d, %d)\n", j, i)
					matchCount++
				}
			}
		}
	}

	return matchCount
}

func Part2() int {
	lines := utils.ReadInput("day4")

	gridSize := 3
	matchCount := 0

	for i, line := range lines {
		for j, _ := range line {
			if i+(gridSize-1) < len(lines) && j+(gridSize-1) < len(line) {
				var minigrid []string

				for gridRows := 0; gridRows < gridSize; gridRows++ {
					minigrid = append(minigrid, lines[i+gridRows][j:j+gridSize])
				}

				if hasXmas(minigrid) {
					fmt.Printf("Found X centred on: (%d, %d)\n", j+1, i+1)
					matchCount++
				}
			}
		}
	}

	return matchCount
}

func hasXmas(grid []string) bool {
	word := "MAS"
	lenOfWord := len(word)

	matchCount := 0

	for i, line := range grid {
		for j, _ := range line {
			if i+(lenOfWord-1) < len(grid) && j+(lenOfWord-1) < len(line) {
				// top to bottom, left to right
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(grid[i+wordParts][j+wordParts])
				}

				if foundWord == word {
					fmt.Printf("Found diagonal (t2b, l2r) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if i-(lenOfWord-1) >= 0 && j+(lenOfWord-1) < len(line) {
				// bottom to top, left to right
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(grid[i-wordParts][j+wordParts])
				}

				if foundWord == word {
					fmt.Printf("Found diagonal (b2t,l2r) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if i+(lenOfWord-1) < len(grid) && j-(lenOfWord-1) >= 0 {
				// top to bottom, right to left
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(grid[i+wordParts][j-wordParts])
				}

				if foundWord == word {
					fmt.Printf("Found diagonal (t2b,r2l) at (%d, %d)\n", j, i)
					matchCount++
				}
			}

			if i-(lenOfWord-1) >= 0 && j-(lenOfWord-1) >= 0 {
				// bottom to top, right to left
				foundWord := ""
				for wordParts := 0; wordParts < lenOfWord; wordParts++ {
					foundWord += string(grid[i-wordParts][j-wordParts])
				}

				if foundWord == word {
					fmt.Printf("Found diagonal (b2t,r2l) at (%d, %d)\n", j, i)
					matchCount++
				}
			}
		}
	}

	return matchCount == 2
}

package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(dayFolder string) []string {
	file, err := os.Open(dayFolder + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Println()
	}
}

package utils

import (
	"bufio"
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

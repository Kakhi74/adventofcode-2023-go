package day1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type Part1 struct{}

func (d Part1) Execute() (int, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory ", err)
		return 0, err
	}

	inputPath := filepath.Join(dir, "day1/input.txt")

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return 0, err
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += d.extractNumber(d.removeNonNumeric(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file ", err)
	}

	fmt.Println("1-1: ", sum)
	return sum, nil
}

func (d Part1) removeNonNumeric(input string) string {
	var builder strings.Builder
	for _, char := range input {
		if unicode.IsDigit(char) {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func (d Part1) extractNumber(numbers string) int {
	if len(numbers) == 0 {
		return 0
	}
	first := string(numbers[0])
	last := string(numbers[len(numbers)-1])
	numStr := first + last
	newNum, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to int ", err)
		return 0
	}
	return newNum
}

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

var numberMap = map[string]rune{
	"zero":  '0',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

type Part2Naive struct{}

func (d Part2Naive) Execute() (int, error) {
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
		sum += d.extractNum(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file ", err)
	}

	fmt.Println("1-2-N: ", sum)
	return sum, nil
}

func (d Part2Naive) extractNum(input string) int {
	fmt.Printf("Input: %s\n", input)
	extractedNumbers := d.wordsToNum(input)
	if len(extractedNumbers) == 0 {
		return 0
	}

	firstNum := string(extractedNumbers[0])
	lastNum := string(extractedNumbers[len(extractedNumbers)-1])
	num, err := strconv.Atoi(firstNum + lastNum)
	if err != nil {
		fmt.Println("Error converting string to int ", err)
		return 0
	}

	fmt.Printf("Extracted numbers: %s\n", extractedNumbers)
	fmt.Printf("result: %d\n", num)

	return num
}

func (d Part2Naive) wordsToNum(s string) string {
	var numericBuilder strings.Builder
	var currentWord strings.Builder

	for _, char := range s {
		if unicode.IsLetter(char) {
			currentWord.WriteRune(char)
			if currentWord.Len() > 2 {
				for key, value := range numberMap {
					currentWordStr := currentWord.String()
					index := strings.Index(currentWordStr, key)
					if index != -1 {
						numericBuilder.WriteRune(value)
						updatedWord := currentWordStr[index+len(key)-1:]
						currentWord.Reset()
						currentWord.WriteString(updatedWord)
						break
					}
				}
			}
		} else if unicode.IsDigit(char) {
			numericBuilder.WriteRune(char)
			currentWord.Reset()
		}
	}
	return numericBuilder.String()
}

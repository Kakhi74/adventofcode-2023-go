package day1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var numbersMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type Part2Performant struct{}

func (d Part2Performant) Execute() (int, error) {
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
		sum += d.stringToNum(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file ", err)
	}

	fmt.Println("1-2-P: ", sum)
	return sum, nil
}

func (d Part2Performant) stringToNum(s string) int {
	indexNumMap := make(map[int]string, 0)
	for key, value := range numbersMap {
		wordIndexes := d.findAllIndexes(s, key)
		for _, index := range wordIndexes {
			indexNumMap[index] = value
		}
		numIndexes := d.findAllIndexes(s, value)
		for _, index := range numIndexes {
			indexNumMap[index] = value
		}
	}
	if len(indexNumMap) < 1 {
		return 0
	}
	keys := make([]int, 0, len(indexNumMap))
	for key := range indexNumMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	first := indexNumMap[keys[0]]
	last := indexNumMap[keys[len(keys)-1]]
	num, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Println("Error converting string to int ", err)
		return 0
	}
	return num
}

func (d Part2Performant) findAllIndexes(s, substr string) []int {
	var indexes []int
	start := 0
	for {
		index := strings.Index(s[start:], substr)
		if index == -1 {
			break
		}
		actualIndex := start + index
		indexes = append(indexes, actualIndex)
		start = actualIndex + 1
	}
	return indexes
}

package day2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var bagContent = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Part1 struct{}

func (d Part1) Execute() (int, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory ", err)
		return 0, err
	}

	inputPath := filepath.Join(dir, "day2/input.txt")

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return 0, err
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += d.possibleGameNumber(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file ", err)
	}

	fmt.Println("2-1: ", sum)
	return sum, nil
}

func (d Part1) possibleGameNumber(line string) int {
	game := strings.Split(line, ":")
	gameSets := strings.Split(game[1], ";")
	for _, set := range gameSets {
		cubes := strings.Split(set, ",")
		for _, c := range cubes {
			cube := strings.Split(strings.TrimSpace(c), " ")
			qty, err := strconv.Atoi(cube[0])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return 0
			}
			color := cube[1]
			maxQty, ok := bagContent[color]
			if !ok {
				return 0
			}
			if qty > maxQty {
				return 0
			}
		}
	}
	return d.getGameNumber(game[0])
}

func (d Part1) getGameNumber(s string) int {
	gameNumberStr := strings.Split(s, " ")[1]
	gameNumber, err := strconv.Atoi(gameNumberStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return gameNumber
}

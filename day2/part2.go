package day2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Part2 struct{}

func (d Part2) Execute() (int, error) {
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
		sum += d.powerOfFewestNumCubesToMakeGamePossible(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file ", err)
	}

	fmt.Println("2-2 ", sum)
	return sum, nil
}

func (d Part2) powerOfFewestNumCubesToMakeGamePossible(line string) int {
	blueQty := 0
	redQty := 0
	greenQty := 0
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
			switch color {
			case "blue":
				if qty > blueQty {
					blueQty = qty
				}
			case "red":
				if qty > redQty {
					redQty = qty
				}
			case "green":
				if qty > greenQty {
					greenQty = qty
				}
			}
		}
	}
	return blueQty * redQty * greenQty
}

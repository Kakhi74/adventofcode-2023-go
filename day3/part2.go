package day3

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type Part2 struct{}

func (d Part2) Execute() (int, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory ", err)
		return 0, err
	}

	inputPath := filepath.Join(dir, "day3/input.txt")

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 0)

	for scanner.Scan() {
		row := d.createRow(scanner.Text())
		matrix = append(matrix, row)
	}

	sum := d.getSum(matrix)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file ", err)
	}

	fmt.Println("3-2: ", sum)
	return sum, nil
}

func (d Part2) createRow(line string) []rune {
	row := make([]rune, 0)
	for _, c := range line {
		if c == '*' {
			row = append(row, 'S')
		} else if unicode.IsDigit(c) {
			row = append(row, c)
		} else {
			row = append(row, 'D')
		}
	}
	return row
}

func (d Part2) getSum(matrix [][]rune) int {
	numbersMap := make(map[string][]int)
	for i := 0; i < len(matrix); i++ {
		numberBuilder := strings.Builder{}
		for j := 0; j < len(matrix[i]); j++ {
			if unicode.IsDigit(matrix[i][j]) {
				numberBuilder.WriteString(string(matrix[i][j]))
				if j != len(matrix[i])-1 {
					continue
				}
			}
			foundSpecial := false
			numberlen := numberBuilder.Len()
			if numberlen == 0 {
				continue
			}
			for x := i - 1; x <= i+1; x++ {
				if x < 0 || x > len(matrix)-1 {
					continue
				}
				for y := j - numberlen - 1; y <= j; y++ {
					if y < 0 || y > len(matrix[x])-1 {
						continue
					}
					if matrix[x][y] == 'S' {
						number, err := strconv.Atoi(numberBuilder.String())
						if err != nil {
							fmt.Println("Error converting string to int: ", numberBuilder.String(), " - error: ", err)
							number = 1
						}
						special_key := fmt.Sprintf("%d-%d", x, y)
						numbersMap[special_key] = append(numbersMap[special_key], number)
						foundSpecial = true
						break
					}
				}
				if foundSpecial {
					break
				}
			}
			numberBuilder.Reset()
		}
	}

	sum := 0
	for _, numbers := range numbersMap {
		multiplier := 1
		if len(numbers) < 2 {
			continue
		}
		for _, number := range numbers {
			multiplier *= number
		}
		sum += multiplier
	}

	return sum
}

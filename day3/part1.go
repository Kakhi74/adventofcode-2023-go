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

type Part1 struct{}

func (d Part1) Execute() (int, error) {
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

	fmt.Println("3-1: ", sum)
	return sum, nil
}

func (d Part1) createRow(line string) []rune {
	row := make([]rune, 0)
	for _, c := range line {
		if c == '.' {
			row = append(row, 'D')
		} else if isSpecial(c) {
			row = append(row, 'S')
		} else if unicode.IsDigit(c) {
			row = append(row, c)
		}
	}
	return row
}

func isSpecial(r rune) bool {
	return !unicode.IsDigit(r) && !(r == '.')
}

func (d Part1) getSum(matrix [][]rune) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		numberBuilder := strings.Builder{}
		for j := 0; j < len(matrix[i]); j++ {
			if unicode.IsDigit(matrix[i][j]) {
				numberBuilder.WriteString(string(matrix[i][j]))
			} else if !unicode.IsDigit(matrix[i][j]) || j == len(matrix[i])-1 {
				foundSpecial := false
				numberlen := numberBuilder.Len()
				if numberlen == 0 {
					continue
				}
				for x := i - 1; x <= i+1; x++ {
					if x < 0 || x > len(matrix)-1 {
						continue
					}
					for y := j - numberlen - 2; y <= j; y++ {
						if y < 0 || y > len(matrix[x])-1 {
							continue
						}
						if matrix[x][y] == 'S' {
							number, err := strconv.Atoi(numberBuilder.String())
							if err == nil {
								fmt.Println("matrix[", x, "][", y, "] = ", number)
								sum += number
							}
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
	}
	return sum
}

package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type PartExecutor interface {
	Execute() (int, error)
}

type Day struct {
	Part1 PartExecutor
	Part2 PartExecutor
}

func main() {
	executors := map[string]Day{
		"1": {Part1: &day1.Part1{}, Part2: &day1.Part2Performant{}},
		"2": {Part1: &day2.Part1{}, Part2: &day2.Part2{}},
	}

	if len(os.Args) < 2 {
		executeLatestDay(executors)
		return
	}

	if len(os.Args) == 3 && os.Args[1] == "-d" {
		day, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid day:", os.Args[2])
			os.Exit(1)
		}
		executeFullDay(executors, fmt.Sprintf("%d", day))
		return
	}

	if len(os.Args) == 5 && os.Args[1] == "-d" && os.Args[3] == "-p" {
		day, err1 := strconv.Atoi(os.Args[2])
		part, err2 := strconv.Atoi(os.Args[4])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid day or part:", os.Args[2], os.Args[4])
			os.Exit(1)
		}
		executeSpecificPart(executors, fmt.Sprintf("%d-%d", day, part))
		return
	}

	fmt.Println("Usage: go run main.go [-d day [-p part]]")
	os.Exit(1)
}

func executeLatestDay(executors map[string]Day) {
	keys := make([]string, 0, len(executors))
	for key := range executors {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	latestDay := executors[keys[len(keys)-1]]
	executeDay(latestDay)
}

func executeFullDay(executors map[string]Day, day string) {
	if dayExecutor, exists := executors[day]; exists {
		executeDay(dayExecutor)
	} else {
		fmt.Printf("No executor found for day %s\n", day)
		os.Exit(1)
	}
}

func executeSpecificPart(executors map[string]Day, dayPart string) {
	day := dayPart[:len(dayPart)-2]
	part := dayPart[len(dayPart)-1:]
	if dayExecutor, exists := executors[day]; exists {
		if part == "1" {
			executePart(dayExecutor.Part1)
		} else if part == "2" {
			executePart(dayExecutor.Part2)
		} else {
			fmt.Printf("No part %s found for day %s\n", part, day)
			os.Exit(1)
		}
	} else {
		fmt.Printf("No executor found for day-part %s\n", dayPart)
		os.Exit(1)
	}
}

func executeDay(day Day) {
	executePart(day.Part1)
	executePart(day.Part2)
}

func executePart(part PartExecutor) {
	_, err := part.Execute()
	if err != nil {
		fmt.Println("Error executing:", err)
		os.Exit(1)
	}
}

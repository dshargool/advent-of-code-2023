package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputToLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	//fmt.Println(TestText() == 8)
	//fmt.Println(Problem1())
	fmt.Println(TestTextTwo() == 2286)

	fmt.Println(Problem2())
}

type bag struct {
	red   int
	green int
	blue  int
}

func (b bag) calcPower() int {
	return b.red * b.blue * b.green
}

func parseGamePossible(line string, bagMax bag) int {
	possible := true
	last_count := 99
	line = line[5:]
	contents := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(contents[0])
	pulls := strings.Split(contents[1], ";")
	for _, pull := range pulls {
		cubes := strings.Split(pull, ", ")
		for _, cube := range cubes {
			colours := strings.Fields(cube)
			for _, colour := range colours {
				if strings.Contains(colour, "red") {
					possible = possible && (last_count <= bagMax.red)
				} else if strings.Contains(colour, "green") {
					possible = possible && (last_count <= bagMax.green)
				} else if strings.Contains(colour, "blue") {
					possible = possible && (last_count <= bagMax.blue)
				} else {
					last_count, _ = strconv.Atoi(colour)
				}
			}
		}
	}
	if possible {
		return gameId
	} else {
		return 0
	}
}

func parseGamePower(line string) int {
	minBag := bag{
		red:   0,
		green: 0,
		blue:  0,
	}
	last_count := 99
	line = line[5:]
	contents := strings.Split(line, ": ")
	pulls := strings.Split(contents[1], ";")
	for _, pull := range pulls {
		cubes := strings.Split(pull, ", ")
		for _, cube := range cubes {
			colours := strings.Fields(cube)
			for _, colour := range colours {
				if strings.Contains(colour, "red") && last_count > minBag.red {
					minBag.red = last_count
				} else if strings.Contains(colour, "green") && last_count > minBag.green {
					minBag.green = last_count
				} else if strings.Contains(colour, "blue") && last_count > minBag.blue {
					minBag.blue = last_count
				} else {
					last_count, _ = strconv.Atoi(colour)
				}
			}
		}
	}
	power := minBag.calcPower()
	fmt.Println(pulls, power, minBag)
	return power
}

func TestText() int {
	var sum int
	bag := bag{
		red:   12,
		green: 13,
		blue:  14,
	}
	lines := ReadInputToLines("./test.txt")
	for _, line := range lines {
		sum += parseGamePossible(line, bag)
	}
	return sum
}

func Problem1() int {
	var sum int
	bag := bag{
		red:   12,
		green: 13,
		blue:  14,
	}
	lines := ReadInputToLines("./p1.txt")
	for _, line := range lines {
		sum += parseGamePossible(line, bag)
	}
	return sum
}

func TestTextTwo() int {
	var sum int
	lines := ReadInputToLines("./test.txt")
	for _, line := range lines {
		sum += parseGamePower(line)
	}
	return sum
}

func Problem2() int {
	var sum int
	lines := ReadInputToLines("./p1.txt")
	for _, line := range lines {
		sum += parseGamePower(line)
	}
	return sum
}

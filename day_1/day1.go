package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(TestText() == 142)

	fmt.Println(Problem1() == 54940)

	fmt.Println(TestTextTwo() == 281)

	fmt.Println(Problem2() == 54208)
}

func ReadInputToLines(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Couldn't read test file")
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func StripAlpha(input string) string {
	var result strings.Builder
	for _, char := range input {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') {
			result.WriteRune(char)
		}
	}

	return result.String()
}

type StringDigitLoc struct {
	index int
	value int
}

func findStringDigits(input string) int {
	// we include 0 so that we can use modulo easily below
	digitStr := []string{
		"0,", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	digits := make([]StringDigitLoc, 0)

	for i, digit := range digitStr {
		firstInd := strings.Index(input, digit)
		lastInd := strings.LastIndex(input, digit)

		if firstInd != lastInd && lastInd != -1 {
			digits = append(digits, StringDigitLoc{
				index: lastInd,
				value: i % 10,
			})
		}
		if firstInd != -1 {
			digits = append(digits, StringDigitLoc{
				index: firstInd,
				value: i % 10,
			})
		}
	}
	if len(digits) > 0 {
		sort.Slice(digits, func(i, j int) bool {
			return digits[i].index < digits[j].index
		})
		value := (digits[0].value * 10) + digits[len(digits)-1].value
		return value
	}
	return 0
}

func TestText() int {
	var sum int64
	lines := ReadInputToLines("./test.txt")
	for _, line := range lines {
		line := StripAlpha(line)
		num, _ := strconv.ParseInt(line, 10, 0)
		right := num % 10
		for num > 10 {
			num /= 10
		}
		sum += num*10 + right
	}
	return int(sum)
}

func Problem1() int {
	var sum int
	lines := ReadInputToLines("./p1.txt")
	for _, line := range lines {
		line := StripAlpha(line)
		num, _ := strconv.ParseInt(line, 10, 0)
		right := num % 10
		for num > 10 {
			num /= 10
		}
		sum += int(num*10 + right)
	}
	return sum
}

func TestTextTwo() int {
	var sum int
	lines := ReadInputToLines("./test2.txt")
	for _, line := range lines {
		sum += findStringDigits(line)
	}
	return sum
}

func Problem2() int {
	var sum int
	lines := ReadInputToLines("./p1.txt")
	for _, line := range lines {
		sum += findStringDigits(line)
	}
	return sum
}

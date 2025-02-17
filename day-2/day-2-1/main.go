package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var inputs [][]int

	// Read file line by line using bufio.Scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			row = append(row, num)
		}
		inputs = append(inputs, row)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	part1, part2 := 0, 0

	// Analyze each report and check if it's safe
	for _, row := range inputs {
		if isSafe(row) {
			part1++
			part2++
			continue

		}
		// if canBeMadeSafeByRemovingOneLevel(row) {
		// 	part2++
		// }

		for _, alternateReport := range generateAlternateReports(row) {
			if isSafe(alternateReport) {
				part2++
				break
			}
		}
	}

	// Output the result
	fmt.Println("The result is:", part1, part2)
}

// Checks if the report is safe based on the rules
func isSafe(elems []int) bool {
	prev := elems[0]
	increase := true
	if elems[1] == elems[0] {
		return false
	}
	if elems[1] < elems[0] {
		increase = false
	}
	for i := 1; i < len(elems); i++ {
		cur := elems[i]
		if increase && cur < prev {
			return false
		}
		if !increase && cur > prev {
			return false
		}
		diff := absInt(prev - cur)
		if diff < 1 || diff > 3 {
			return false
		}
		prev = cur
	}
	return true
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func generateAlternateReports(row []int) [][]int {
	var result [][]int
	for i := 0; i < len(row); i++ {
		newSlice := append([]int{}, row[:i]...)
		newSlice = append(newSlice, row[i+1:]...)

		result = append(result, append(row[:i], row[i+1:]...))
	}
	return result
}

// Tries to make the report safe by removing one level
func canBeMadeSafeByRemovingOneLevel(row []int) bool {
	// Try removing each level and check if the resulting report is safe
	for i := 0; i < len(row); i++ {
		// Create a new report by removing the ith level
		newRow := append(row[:i], row[i+1:]...)
		if isSafe(newRow) {
			return true
		}
	}
	return false
}
